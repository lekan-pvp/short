package memrepo

import "context"

// SoftDelete is a plug for implementation the function in DBRepo type.
// The function is not used.
func (r MemoryRepo) SoftDelete(_ context.Context, _ []string, _ string) error {
	return nil
}
