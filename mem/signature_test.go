package mem_test

import (
	"testing"

	"github.com/jaimelopez/tril3ro/mem"
)

func TestGeneratePattern(t *testing.T) {
	tests := []struct {
		name    string
		combo   string
		pattern string
		mask    string
	}{
		{
			"Single asterix",
			"* C2 85 C0 7E * 8B D0 E8",
			"? C2 85 C0 7E ? 8B D0 E8",
			"?xxxx?xxx",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pattern, mask := mem.GeneratePattern(test.combo)

			if pattern != test.pattern {
				t.Errorf("wrong pattern, expected %s, got %s", test.pattern, pattern)
			}

			if mask != test.mask {
				t.Errorf("wrong mask, expected %s, got %s", test.mask, mask)
			}
		})
	}
}
