package file

import "errors"

var (
	// ErrNotReferencedObject it's raised when a non pointer is received trying to parse into a struct
	ErrNotReferencedObject = errors.New("object not a pointer or reference")

	// ErrNotReferencedObject it's raised when a non pointer is received trying to parse into a struct
	ErrInvalidObject = errors.New("object specified is not valid")

	// ErrNonStructObject it's raised when the underlaying object is not an struct
	ErrNonStructObject = errors.New("underlaying object is not a struct")
)
