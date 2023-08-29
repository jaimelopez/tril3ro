package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/common"
	"golang.org/x/sys/unix"
)

// Read certain memory address
func (r *reader[T]) ReadOf(addr common.Addr, into *T, size uint) error {
	buffer := (*byte)(unsafe.Pointer(into))
	sz := int(size)

	n, err := unix.ProcessVMReadv(
		int(r.processID),
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
