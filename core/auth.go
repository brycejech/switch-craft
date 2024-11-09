package core

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"switchcraft/types"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("incorrect hash format")
	ErrIncompatibleVersion = errors.New("incompatible argon2 version")
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultHashParams = &hashParams{
	memory:      64 * 1024,
	iterations:  6,
	parallelism: 4,
	saltLength:  16,
	keyLength:   32,
}

func (c *Core) Authn(ctx context.Context, username string, password string) (account *types.Account, ok bool) {
	var (
		passwordsMatch bool
		err            error
	)

	if account, err = c.AccountGetOne(ctx,
		c.NewAccountGetOneArgs(nil, nil, nil, &username),
	); err != nil {
		return nil, false
	}

	if account.Password == nil {
		return nil, false
	}

	if passwordsMatch, err = c.AuthPasswordCheck(
		password,
		*account.Password,
	); err != nil {
		return nil, false
	}

	if !passwordsMatch {
		return nil, false
	}

	return account, true
}

func (c *Core) AuthCreateSigningKey(bitLength uint32) (string, error) {
	if bitLength < 256 {
		return "", errors.New("error: bitLength must be >= 256")
	}
	if bitLength%8 != 0 {
		return "", errors.New("error: bitLength must be divisible by 8")
	}

	bytes, err := randomBytes(bitLength / 4)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (c *Core) AuthPasswordHash(password string) (encodedHash string, err error) {
	salt, err := randomBytes(defaultHashParams.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(
		[]byte(password),
		salt,
		defaultHashParams.iterations,
		defaultHashParams.memory,
		defaultHashParams.parallelism,
		defaultHashParams.keyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	tmpHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		defaultHashParams.memory,
		defaultHashParams.iterations,
		defaultHashParams.parallelism,
		b64Salt,
		b64Hash,
	)

	return tmpHash, nil
}

func (c *Core) AuthPasswordCheck(password, encodedHash string) (match bool, err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *hashParams, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &hashParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func randomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
