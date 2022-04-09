package memrepo

import "context"

// PingDB is plug for implementation in DBRepo.
// The function is not used.
func (r *MemoryRepo) PingDB(_ context.Context) error {
	return nil
}
