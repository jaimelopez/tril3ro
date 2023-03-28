package proc

// Reader struct to manage memory reading operations
type Reader[T any] struct {
	*Process
}

// NewReader instantiates a new memory reader for specified data struct
func NewReader[T any](procces *Process) *Reader[T] {
	return &Reader[T]{procces}
}

// Writer struct to manage memory writing operations
type Writer[T any] struct {
	*Process
}

// NewWriter instantiates a new memory writer for specified data struct
func NewWriter[T any](procces *Process) *Writer[T] {
	return &Writer[T]{procces}
}
