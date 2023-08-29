package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/common"
)

// reader struct to manage memory reading operations
type reader[T any] struct {
	*manager
}

// Read certain memory address
func (r *reader[T]) Read(addr common.Addr, into *T) error {
	var et T

	return r.ReadOf(addr, into, uint(unsafe.Sizeof(et)))
}

// NewReader instantiates a new memory reader for specified data struct
func NewReader[T any](opts ...Option) (*reader[T], error) {
	m, err := newManager(opts...)
	if err != nil {
		return nil, err
	}

	return &reader[T]{m}, nil
}

// NewReaderForProc instantiates a new memory reader for specified data struct using the default handler
func NewReaderForProc[T any](processID uint32) (*reader[T], error) {
	return NewReader[T](WithDefaultHandler(processID))
}
