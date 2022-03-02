package dbrepo

import (
	"context"
	"database/sql"
	"github.com/lekan-pvp/short/internal/config"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	UUID          string `json:"uuid" db:"user_id"`
	ShortURL      string `json:"short_url" db:"short_url"`
	OriginalURL   string `json:"original_url" db:"orig_url"`
	CorrelationID string `json:"correlation_id" db:"correlation_id"`
	DeleteFlag    bool   `json:"delete_flag" db:"is_deleted"`
}

var db *sql.DB

func New() {
	var err error
	databaseDSN := config.GetDatabaseURI()
	db, err = sql.Open("postgres", databaseDSN)
	if err != nil {
		log.Printf("dtatbase connecting error %s", err)
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

func PostURL(ctx context.Context, rec Storage) error {
	log.Println("IN InsertUserRepo short url =", rec.ShortURL)
	_, err := db.ExecContext(ctx, `INSERT INTO users(user_id, short_url, orig_url) VALUES ($1, $2, $3);`, rec.UUID, rec.ShortURL, rec.OriginalURL)

	if err != nil {
		return err
	}

	return nil
}

type OriginURL struct {
	URL     string
	Deleted bool
}

func (u OriginURL) IsDeleted() bool {
	return u.Deleted
}

func GetOriginal(ctx context.Context, short string) (*OriginURL, error) {
	result := &OriginURL{}

	err := db.QueryRowContext(ctx, `SELECT orig_url, is_deleted FROM users WHERE short_url=$1;`, short).Scan(result.URL, result.Deleted)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type ListResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func GetURLsList(ctx context.Context, uuid string) ([]ListResponse, error) {
	var list []ListResponse

	rows, err := db.QueryContext(ctx, `SELECT short_url, orig_url FROM users WHERE user_id=$1`, uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var v ListResponse
		err = rows.Scan(&v.ShortURL, &v.OriginalURL)
		if err != nil {
			return nil, err
		}
		list = append(list, v)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return list, nil
}
