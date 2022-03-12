package memrepo

import "context"

func (r *MemoryRepo) PingDB(_ context.Context) error {
	return nil
}
