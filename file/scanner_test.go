package file_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/jaimelopez/tril3ro/file"
	"golang.org/x/exp/slices"
)

func TestScanBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []string
	}{
		{
			"Simple buffer",
			[]byte("uno"),
			[]string{"uno"},
		},
		{
			"Simple buffer with extra break line",
			[]byte("uno\n"),
			[]string{"uno"},
		},
		{
			"Buffer with two blocks of one element each",
			[]byte("uno\n\ndos"),
			[]string{"uno", "dos"},
		},
		{
			"Buffer with two blocks of two element each",
			[]byte("uno\ndos\n\ntres\ncuatro"),
			[]string{"uno\ndos", "tres\ncuatro"},
		},
		{
			"Buffer with an extra blank line at the beginning",
			[]byte("\n\nuno\ndos\n\ntres\ncuatro"),
			[]string{"uno\ndos", "tres\ncuatro"},
		},
		{
			"Buffer with an extra blank block at the end",
			[]byte("uno\ndos\n\ntres\ncuatro\n\n\n"),
			[]string{"uno\ndos", "tres\ncuatro"},
		},
		{
			"Buffer with an empty block between other blocks",
			[]byte("uno\ndos\n\n\ntres\ncuatro"),
			[]string{"uno\ndos", "tres\ncuatro"},
		},
		{
			"Buffer with whitespaces",
			[]byte("uno\ndos\n  \n\ntres\ncuatro"),
			[]string{"uno\ndos", "tres\ncuatro"},
		},
		{
			"Buffer with tabs",
			[]byte("uno\tdos\ntres\tcuatro"),
			[]string{"uno\tdos\ntres\tcuatro"},
		},
	}

	// It's being tested together with a bufio scanner so technically speaking, this is not a pure unit tests
	// but more a integration test instead.
	// The reason to do so is because the method was introduced to work together with a buffer scanner
	// so I really want to ensure it works properly together with a bufio scanner.
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buffer := bytes.NewBuffer(test.input)
			scan := bufio.NewScanner(buffer)
			scan.Split(file.ScanBlocks)

			blocks := []string{}

			for scan.Scan() {
				blocks = append(blocks, scan.Text())
			}

			if !slices.Equal(blocks, test.expected) {
				t.Error("expected:", test.expected, "got:", blocks)
			}
		})
	}

}
