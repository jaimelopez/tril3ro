package mem_test

import (
	"math/rand"
	"testing"

	"github.com/jaimelopez/tril3ro/mem"
)

func BenchmarkSearchPattern(b *testing.B) {
	total := 10 * 1024
	maxPatternLenght := 20

	data := []byte{}

	for i := 0; i < total; i++ {
		data = append(data, byte(i))
	}

	patterns := [][]byte{}

	for n := 0; n < b.N; n++ {
		start := rand.Intn(len(data))
		max := len(data) - (start + 1)

		if max > maxPatternLenght {
			max = maxPatternLenght
		}

		if max == 0 {
			continue
		}

		end := rand.Intn(max)
		pattern := data[start : start+end]
		patterns = append(patterns, pattern)
	}

	b.ResetTimer()

	for _, pattern := range patterns {
		_, _ = mem.SearchPattern(data, pattern)
	}
}
