package dbrepo

import "context"

// Ping method checks database connection.
// Used in dbhandlers.PingDB handler.
func (r *DBRepo) PingDB(ctx context.Context) error {
	if err := r.db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
