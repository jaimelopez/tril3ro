package mem

import "errors"

var (
	// ErrInvalidPattern represents an error when the specified pattern is empty or invalid
	ErrInvalidPattern = errors.New("invalid pattern")

	// ErrPatternNotFound represents an error when the pattern is not found in the specified data
	ErrPatternNotFound = errors.New("pattern not found")

	// ErrLenghtMismatching is raised when there is a difference in lenght between the signature and the mask
	ErrLenghtMismatching = errors.New("different lengths between signature and mask")

	// ErrWrongTotalBytes represents an error reading or writing the spcified number of bytes
	ErrWrongTotalBytes = errors.New("couldn't perform operation with as many bytes as expected")

	// ErrUnexpectedConversion represents generic unexpected error
	ErrUnexpectedResult = errors.New("unexpected result error")

	// ErrHandlerNotSpecified when handler is nil or not specified
	ErrHandlerNotSpecified = errors.New("nil handler or not specified")
)
