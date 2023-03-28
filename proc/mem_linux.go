package proc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// Read certain memory address
func (r *Reader[T]) Read(addr Addr) (*T, error) {
	var data T

	size := int(unsafe.Sizeof(data))
	buffer := (*byte)(unsafe.Pointer(&data))

	n, err := unix.ProcessVMReadv(
		int(r.ID),
		[]unix.Iovec{{Base: buffer, Len: uint64(size)}},
		[]unix.RemoteIovec{{Base: addr, Len: size}},
		0,
	)

	if err != nil {
		return nil, err
	}

	if n != size {
		return nil, ErrWrongTotalBytes
	}

	return &data, nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr Addr, data T) error {
	size := int(unsafe.Sizeof(data))
	buffer := (*byte)(unsafe.Pointer(&data))

	n, err := unix.ProcessVMWritev(
		int(r.ID),
		[]unix.Iovec{{Base: buffer, Len: uint64(size)}},
		[]unix.RemoteIovec{{Base: addr, Len: size}},
		0,
	)

	if err != nil {
		return err
	}

	if n != size {
		return ErrWrongTotalBytes
	}

	return nil
}
