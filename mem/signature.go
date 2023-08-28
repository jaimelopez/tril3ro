package mem

import (
	"encoding/hex"
	"strings"

	"golang.org/x/exp/slices"
)

func GeneratePattern(signature string, mask string) ([]byte, error) {
	sign := strings.Split(signature, "\\x")

	if len(sign) != len(mask) {
		return nil, ErrLenghtMismatching
	}

	pattern := []byte{}
	wildcards := []byte{'x', '?'}

	for i, val := range strings.ToLower(mask) {
		if slices.Contains(wildcards, byte(val)) {
			pattern = append(pattern, patternWildcard)
			continue
		}

		c, err := hex.DecodeString(string(sign[i]))
		if err != nil {
			return nil, err
		}

		pattern = append(pattern, c...)
	}

	return pattern, nil
}
