package proc

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// Read certain memory address
func (r *Reader[T]) ReadOf(addr Addr, into *T, size uint) error {
	_ = r.open()

	buffer := (*[]byte)(unsafe.Pointer(into))

	if err := windows.ReadProcessMemory(
		r.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(buffer)),
		uintptr(size),
		nil,
	); err != nil {
		return err
	}

	return nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) WriteOf(addr Addr, data T, size uint) error {
	_ = r.open()

	dtw := (*[]byte)(unsafe.Pointer(&data))

	if err := windows.WriteProcessMemory(
		r.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(dtw)),
		uintptr(size),
		nil,
	); err != nil {
		return err
	}

	return nil
}
