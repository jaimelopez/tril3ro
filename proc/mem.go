package proc

import "unsafe"

// Reader struct to manage memory reading operations
type Reader[T any] struct {
	*Process
}

// Read certain memory address
func (r *Reader[T]) Read(addr Addr, into *T) error {
	var et T

	return r.ReadOf(addr, into, uint(unsafe.Sizeof(et)))
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

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr Addr, data T) error {
	var et T

	return r.WriteOf(addr, data, uint(unsafe.Sizeof(et)))
}
