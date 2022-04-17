package dbrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/models"
)

func (r *DBRepo) GetStats(ctx context.Context) (models.Stat, error) {
	var stats models.Stat
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(orig_url), COUNT(DISTINCT user_id) FROM users`).Scan(&stats.URLs, &stats.Users)
	if err != nil {
		return models.Stat{}, err
	}
	return stats, nil
}
