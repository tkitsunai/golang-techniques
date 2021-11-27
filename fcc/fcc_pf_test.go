package fcc_test

import (
	"fmt"
	"golang-techniques/fcc"
	"testing"
)

func BenchmarkFCC(b *testing.B) {

	b.Run("memory allocation once", func(b *testing.B) {
		base := make(fcc.Users, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			base[i] = fcc.User{
				ID:   fmt.Sprintf("%d", i),
				Name: fmt.Sprintf("NAME%d", i),
			}
		}
	})
}
