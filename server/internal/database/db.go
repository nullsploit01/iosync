package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/env"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

func NewDbClient(ctx context.Context) (*ent.Client, error) {
	databaseUser := env.GetString("DATABASE_USER", "postgres")
	databasePassword := env.GetString("DATABASE_PASSWORD", "password")
	databaseHost := env.GetString("DATABASE_HOST", "localhost")
	databasePort := env.GetString("DATABASE_PORT", "5432")
	databaseName := env.GetString("DATABASE_NAME", "iosync")

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	poolConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database url: %w", err)
	}
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db := stdlib.OpenDBFromPool(pool)
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed to create schema resources: %w", err)
	}

	log.Println("connected to database server", databaseHost)
	return client, nil
}
