package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/lekan-pvp/short/internal/config"
	"github.com/lekan-pvp/short/internal/makeshort"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"log"
)

type Storage struct {
	UUID          string `json:"uuid" db:"user_id"`
	ShortURL      string `json:"short_url" db:"short_url"`
	OriginalURL   string `json:"original_url" db:"orig_url"`
	CorrelationID string `json:"correlation_id" db:"correlation_id"`
	DeleteFlag    bool   `json:"delete_flag" db:"is_deleted"`
}

type URL struct {
	URL string `json:"url"`
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

func PostURL(ctx context.Context, rec Storage) (string, error) {
	log.Println("IN InsertUserRepo short url =", rec.ShortURL)
	_, err := db.ExecContext(ctx, `INSERT INTO users(user_id, short_url, orig_url) VALUES ($1, $2, $3);`, rec.UUID, rec.ShortURL, rec.OriginalURL)

	var result string

	if err != nil {
		if err.(*pq.Error).Code == pgerrcode.UniqueViolation {
			notOk := db.QueryRowContext(ctx, `SELECT short_url FROM users WHERE orig_url=$1 AND user_id=$2;`, rec.OriginalURL, rec.UUID).Scan(&result)
			if notOk != nil {
				return "", notOk
			}
			return "", err
		}
	}

	return result, nil
}

type OriginURL struct {
	URL     string
	Deleted bool
}

func (u OriginURL) IsDeleted() bool {
	return u.Deleted
}

func GetOriginal(ctx context.Context, short string) (*OriginURL, error) {
	log.Println("GetOriginal IN DB")
	result := &OriginURL{}

	err := db.QueryRowContext(ctx, `SELECT orig_url, is_deleted FROM users WHERE short_url=$1;`, short).Scan(&result.URL, &result.Deleted)
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

type BatchResponse struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type BatchRequest struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}

func BatchShorten(ctx context.Context, uuid string, in []BatchRequest) ([]BatchResponse, error) {
	var res []BatchResponse
	base := config.GetBaseURL()

	log.Printf("in is %v", in)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users(user_id, short_url, orig_url, correlation_id) 
												VALUES($1, $2, $3, $4)`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	for _, v := range in {
		log.Printf("%s %s", v.OriginalURL, v.CorrelationID)
		short := makeshort.GenerateShortLink(v.OriginalURL, v.CorrelationID)
		log.Printf("short is %s", short)
		if _, err = stmt.ExecContext(ctx, uuid, short, v.OriginalURL, v.CorrelationID); err != nil {
			log.Println("insert exec error")
			return nil, err
		}
		res = append(res, BatchResponse{CorrelationID: v.CorrelationID, ShortURL: base + "/" + short})
	}
	if err := tx.Commit(); err != nil {
		log.Println("commit error")
		return nil, err
	}
	return res, nil
}

type Query struct {
	UserID   string
	ShortURL string
}

func fanOut(input []Query, n int) []chan Query {
	chs := make([]chan Query, 0, n)
	for i, val := range input {
		ch := make(chan Query, 1)
		ch <- val
		chs = append(chs, ch)
		close(chs[i])
	}
	return chs
}

func newWorker(ctx context.Context, stmt *sql.Stmt, tx *sql.Tx, jobs <-chan Query) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for id := range jobs {
		if _, err := stmt.ExecContext(ctx, id.ShortURL, id.UserID); err != nil {
			if err = tx.Rollback(); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}

func SoftDelete(ctx context.Context, in []string, uuid string) error {
	n := len(in)

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	if len(in) == 0 {
		return errors.New("the list of URLs is empty")
	}

	var query []Query

	for _, v := range in {
		log.Println(v)
		var q Query
		q.UserID = uuid
		q.ShortURL = v
		query = append(query, q)
	}

	fanOutChs := fanOut(query, n)

	stmt, err := tx.PrepareContext(ctx, `UPDATE users SET is_deleted=TRUE WHERE short_url=$1 AND user_id=$2`)
	if err != nil {
		log.Println("stmt error")
		return err
	}
	defer stmt.Close()

	jobs := make(chan Query, n)

	g, _ := errgroup.WithContext(ctx)

	for i := 1; i <= 3; i++ {
		g.Go(func() error {
			err = newWorker(ctx, stmt, tx, jobs)
			if err != nil {
				log.Println("error in g.Go")
				return err
			}
			return nil
		})
	}

	for _, item := range fanOutChs {
		jobs <- <-item
	}
	close(jobs)

	if err := g.Wait(); err != nil {
		log.Println("Wait error")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println("Commit error")

		return err
	}

	return nil
}
