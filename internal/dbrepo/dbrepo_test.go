package dbrepo

import (
	"context"
	"github.com/lekan-pvp/short/internal/config"
	"log"
	"testing"
)

func BenchmarkPing(b *testing.B) {
	config.New()
	New()
	ctx, _ := context.WithCancel(context.Background())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := Ping(ctx); err != nil {
			log.Println(err)
		}
	}
}
