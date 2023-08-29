package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
	"golang.org/x/sys/unix"
)

// Write certain data into a particular memory address
func (w *writer[T]) WriteOf(addr proc.Addr, data T, size uint) error {
	buffer := (*byte)(unsafe.Pointer(&data))
	sz := int(size)

	n, err := unix.ProcessVMWritev(
		int(w.processID),
		[]unix.Iovec{{Base: buffer, Len: uint64(sz)}},
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
