package dbrepo

import (
	"context"
	"database/sql"
	"github.com/lekan-pvp/short/internal/config"
	"log"
)

var db *sql.DB

func New() {
	var err error
	databaseDSN := config.GetDatabaseURI()
	db, err = sql.Open("postgres", databaseDSN)
	if err != nil {
		log.Fatal("database connecting error", err)
	}

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users(id SERIAL, user_id VARCHAR, short_url VARCHAR NOT NULL, orig_url VARCHAR NOT NULL, correlation_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE, PRIMARY KEY (id), UNIQUE (orig_url));`)
	if err != nil {
		log.Fatal("create table error", err)
	}
}

func Ping(ctx context.Context) error {
	if err := db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
