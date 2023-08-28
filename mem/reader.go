package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

// Reader struct to manage memory reading operations
type Reader[T any] struct {
	*handler
}

// Read certain memory address
func (r *Reader[T]) Read(addr proc.Addr, into *T) error {
	var et T

	return r.ReadOf(addr, into, uint(unsafe.Sizeof(et)))
}

// NewReader instantiates a new memory reader for specified data struct
func NewReader[T any](processID uint32) (*Reader[T], error) {
	h, err := NewHandler(processID)
	if err != nil {
		return nil, err
	}

	return &Reader[T]{h}, nil
}
