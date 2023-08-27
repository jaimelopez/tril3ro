package mem_test

import (
	"testing"

	"github.com/jaimelopez/tril3ro/mem"
)

func TestSearchPattern(t *testing.T) {
	tests := []struct {
		name          string
		data          []byte
		pattern       []byte
		found         bool
		expectedError error
	}{
		{
			"Empty pattern",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{},
			false,
			mem.ErrInvalidPattern,
		},
		{
			"All wildcards",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{0x00, 0x00},
			true,
			nil,
		},
		{
			"Starting with wildcards",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{0x00, 0x22},
			true,
			nil,
		},
		{
			"Ending with wildcards",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{0x24, 0x00},
			true,
			nil,
		},
		{
			"Excceeding lenght wild wildcard",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{0x25, 0x00},
			false,
			mem.ErrPatternNotFound,
		},
		{
			"Not matching at all",
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25},
			[]byte{0x26, 0x00},
			false,
			mem.ErrPatternNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pos, err := mem.SearchPattern(test.data, test.pattern)
			result := (pos >= 0 && err == nil)

			if test.found != result {
				t.Errorf("invalid expectation: expecting %t got %t", test.found, result)
			}

			if err != test.expectedError {
				t.Errorf("invalid expected error: expecting %v got %v", test.expectedError, err)
			}
		})
	}
}
