package dbrepo

import (
	"context"
	"database/sql"
	"github.com/lekan-pvp/short/internal/config"
	"log"
)

type DBRepo struct {
	db *sql.DB
}

// New method for setup database and creating a table.
func New(cfg config.Config) DBRepo {
	var err error
	var r DBRepo

	r.db, err = sql.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		log.Printf("dtatbase connecting error %s", err)
		log.Fatal("database connecting error", err)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	_, err = r.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(id SERIAL, user_id VARCHAR, short_url VARCHAR NOT NULL, orig_url VARCHAR NOT NULL, correlation_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE, PRIMARY KEY (id), UNIQUE (orig_url));`)
	if err != nil {
		log.Fatal("create table error", err)
	}

	return r
}
