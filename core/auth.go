package core

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"switchcraft/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("incorrect hash format")
	ErrIncompatibleVersion = errors.New("incompatible argon2 version")
)

const jwtIssuer = "SwitchCraft"
const jwtLifetime = 24 * time.Hour

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

func (c *Core) Authn(ctx context.Context, username string, password string) (*types.Account, bool) {
	var (
		account        *types.Account
		passwordsMatch bool
		err            error
	)

	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, false
	}

	if account, err = c.AccountGetOne(ctx,
		c.NewAccountGetOneArgs(nil, nil, nil, &username),
	); err != nil {
		c.logger.Error(tracer, err.Error(), nil)
		return nil, false
	}

	if account.Password == nil {
		return nil, false
	}

	if passwordsMatch, err = c.AuthPasswordCheck(
		password,
		*account.Password,
	); err != nil {
		c.logger.Error(tracer, err.Error(), nil)
		return nil, false
	}

	if !passwordsMatch {
		c.logger.Error(tracer, "Incorrect username/password combination", nil)
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

	bytes, err := randomBytes(bitLength / 8)
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

func (c *Core) AuthCreateJWT(account *types.Account) (string, error) {
	var (
		token    *jwt.Token
		tokenStr string
		err      error
		key      = c.jwtSigningKey
	)

	if err = validateJWTSigningKey(key); err != nil {
		return "", fmt.Errorf("core.AuthCreateJWT: %w", err)
	}

	token = jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			// Registered claims
			"iss": jwtIssuer,
			"aud": []string{jwtIssuer},
			"sub": account.Username,
			"exp": time.Now().Add(jwtLifetime).Unix(), // 1 day
			"iat": time.Now().Unix(),

			// Custom claims
			"account": account,
		},
	)

	if tokenStr, err = token.SignedString(key); err != nil {
		return "", fmt.Errorf("core.AuthCreateJWT error signing JWT: %w", err)
	}

	return tokenStr, nil
}

func (c *Core) AuthValidateJWT(jwtString string) (*types.Account, error) {
	var (
		account *types.Account
		token   *jwt.Token
		err     error
	)

	if token, err = parseJWT(c.jwtSigningKey, jwtString); err != nil {
		return nil, fmt.Errorf("core.AuthValidateJWT: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("core.AuthValidateJWT error parsing JWT claims")
	}

	issuer, ok := claims["iss"].(string)
	if !ok || issuer != jwtIssuer {
		return nil, errors.New("core.AuthValidateJWT invalid JWT issuer")
	}

	accountMap, ok := claims["account"].(map[string]any)
	if !ok {
		return nil, errors.New("core.AuthValidateJWT error casting claims.account")
	}

	if account, err = mapClaimsAccount(accountMap); err != nil {
		return nil, fmt.Errorf("core.AuthValidateJWT: %w", err)
	}

	return account, nil
}

func parseJWT(signingKey []byte, jwtString string) (*jwt.Token, error) {
	var (
		token *jwt.Token
		err   error
	)

	if err = validateJWTSigningKey(signingKey); err != nil {
		return nil, err
	}

	var (
		validSigningMethods = jwt.WithValidMethods([]string{"HS512"})
		parseTokenCallback  = getTokenParserCallback(signingKey)
	)
	if token, err = jwt.Parse(jwtString, parseTokenCallback, validSigningMethods); err != nil {
		return nil, fmt.Errorf("core.parseJWT invalid token: %w", err)
	}

	return token, nil
}

func getTokenParserCallback(signingKey []byte) func(*jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("core.parseJWT token signing method mismatch")
		}
		return signingKey, nil
	}
}

func mapClaimsAccount(accountMap map[string]any) (*types.Account, error) {
	account := new(types.Account)

	accountBytes, err := json.Marshal(accountMap)
	if err != nil {
		return nil, fmt.Errorf("core.mapClaimsAccount accountMap marshal error: %w", err)
	}

	if err = json.Unmarshal(accountBytes, account); err != nil {
		return nil, fmt.Errorf("core.mapClaimsAccount error parsing accountMap: %w", err)
	}

	return account, nil
}

func validateJWTSigningKey(key []byte) error {
	if len(key) != 64 {
		return errors.New(
			fmt.Sprintf(
				"core.validateJWTSigningKey: invalid JWT signing key length - expected 512 bits, got %v",
				len(key)*8,
			),
		)
	}
	return nil
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
