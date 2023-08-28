package mem

const patternWildcard = byte(0x00)

func SearchPattern(data []byte, pattern []byte) (int, error) {
	if len(pattern) == 0 && len(data) != 0 {
		return -1, ErrInvalidPattern
	}

	for i := range data {
		for x, p := range pattern {
			if len(data) <= i+x {
				break
			}

			if p != 0x00 && data[i+x] != p {
				break
			}

			if x == len(pattern)-1 {
				return i, nil
			}
		}
	}

	return -1, ErrPatternNotFound
}
