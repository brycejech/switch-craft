package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"switchcraft/cmd/cli"
	"switchcraft/core"
	"switchcraft/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbSSLMode  = os.Getenv("DB_SSL_MODE")
	dbMaxConns = os.Getenv("DB_MAX_CONNECTIONS")
)

var globalCtx = context.Background()

func main() {
	db := mustInitDb(globalCtx)

	var (
		repo            = repository.NewRepository(db)
		accountRepo     = repository.NewAccountRepository(db)
		tenantRepo      = repository.NewTenantRepository(db)
		applicationRepo = repository.NewAppRepository(db)
		featureFlagRepo = repository.NewFeatureFlagRepository(db)
	)

	switchcraft := core.NewCore(
		repo,
		accountRepo,
		tenantRepo,
		applicationRepo,
		featureFlagRepo,
	)

	cli.Start(switchcraft)
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
		log.Fatal(fmt.Sprintf("error building database dsn: %s", err))
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal(fmt.Sprintf("error establishing connection pool: %s", err))
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatal(fmt.Sprintf("error pinging database: %s", err))
	}

	return pool
}
