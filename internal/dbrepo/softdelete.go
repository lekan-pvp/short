package dbrepo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lekan-pvp/short/internal/models"
	"golang.org/x/sync/errgroup"
	"log"
)

func fanOut(input []models.Query, n int) []chan models.Query {
	chs := make([]chan models.Query, 0, n)
	for i, val := range input {
		ch := make(chan models.Query, 1)
		ch <- val
		chs = append(chs, ch)
		close(chs[i])
	}
	return chs
}

func newWorker(ctx context.Context, stmt *sql.Stmt, tx *sql.Tx, jobs <-chan models.Query) error {
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

// SoftDelete is a method to set a DeleteFlag in database.
// Used in dbhandlers.SoftDelete handler.
// This in a concurrency method.
func (r DBRepo) SoftDelete(ctx context.Context, in []string, uuid string) error {
	n := len(in)

	tx, err := r.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	if len(in) == 0 {
		return errors.New("the list of URLs is empty")
	}

	var query []models.Query

	for _, v := range in {
		log.Println(v)
		var q models.Query
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

	jobs := make(chan models.Query, n)

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
