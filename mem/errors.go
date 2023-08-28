package mem

import "errors"

var (
	// ErrInvalidPattern represents an error when the specified pattern is empty or invalid
	ErrInvalidPattern = errors.New("invalid pattern")

	// ErrPatternNotFound represents an error when the pattern is not found in the specified data
	ErrPatternNotFound = errors.New("pattern not found")

	// ErrLenghtMismatching is raised when there is a difference in lenght between the signature and the mask
	ErrLenghtMismatching = errors.New("different lengths between signature and mask")
)
