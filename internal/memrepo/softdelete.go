package memrepo

import "context"

func (r *MemoryRepo) SoftDelete(_ context.Context, _ []string, _ string) error {
	return nil
}
