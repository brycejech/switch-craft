package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"switchcraft/cmd/cli"
	"switchcraft/core"
	"switchcraft/repository"
	"switchcraft/types"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbHost        = os.Getenv("DB_HOST")
	dbPort        = os.Getenv("DB_PORT")
	dbUser        = os.Getenv("DB_USER")
	dbPass        = os.Getenv("DB_PASS")
	dbDatabase    = os.Getenv("DB_DATABASE")
	dbSSLMode     = os.Getenv("DB_SSL_MODE")
	dbMaxConns    = os.Getenv("DB_MAX_CONNECTIONS")
	jwtSigningKey = os.Getenv("JWT_SIGNING_KEY")
)

var globalCtx = context.Background()

func main() {
	jwtSigningKeyBytes := mustGetJWTSigningKey(jwtSigningKey)

	var (
		logger            = types.NewLogger(types.LogLevelInfo)
		db                = mustInitDb(globalCtx)
		repo              = repository.NewRepository(logger, db)
		globalAccountRepo = repository.NewGlobalAccountRepository(logger, db)
		orgAccountRepo    = repository.NewOrgAccountRepository(logger, db)
		orgGroupRepo      = repository.NewOrgGroupRepository(logger, db)
		orgRepo           = repository.NewOrgRepository(logger, db)
		applicationRepo   = repository.NewAppRepository(logger, db)
		featureFlagRepo   = repository.NewFeatureFlagRepository(logger, db)
	)

	switchcraft := core.NewCore(
		logger,
		repo,
		globalAccountRepo,
		orgAccountRepo,
		orgGroupRepo,
		orgRepo,
		applicationRepo,
		featureFlagRepo,
		jwtSigningKeyBytes,
	)

	cli.Start(logger, switchcraft)
}

func mustInitDb(ctx context.Context) *pgxpool.Pool {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s pool_max_conns=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbDatabase,
		dbSSLMode,
		dbMaxConns,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(fmt.Errorf("error building database dsn: %w", err))
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(fmt.Errorf("error establishing connection pool: %w", err))
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatal(fmt.Errorf("error pinging database: %w", err))
	}

	return pool
}

func mustGetJWTSigningKey(key string) []byte {
	bytes, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(fmt.Errorf("invalid JWT signing key: %w", err))
	}

	var (
		byteLen = len(bytes)
		bitLen  = byteLen * 8
	)

	if bitLen != 512 {
		fmt.Println(
			"warning: JWT signing key length invalid - expected 512 bits, got",
			bitLen,
		)
	}

	return bytes
}
