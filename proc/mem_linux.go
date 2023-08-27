package proc

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// Read certain memory address
func (r *Reader[T]) ReadOf(addr Addr, into *T, size uint) error {
	buffer := (*byte)(unsafe.Pointer(into))
	sz := int(size)

	n, err := unix.ProcessVMReadv(
		int(r.ID),
		[]unix.Iovec{{Base: buffer, Len: uint64(size)}},
		[]unix.RemoteIovec{{Base: addr, Len: sz}},
		0,
	)

	if err != nil {
		return err
	}

	if n != sz {
		return ErrWrongTotalBytes
	}

	return nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) WriteOf(addr Addr, data T, size uint) error {
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
