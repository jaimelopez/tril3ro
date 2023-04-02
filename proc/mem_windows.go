package proc

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// Read certain memory address
func (r *Reader[T]) Read(addr Addr, into *T) error {
	_ = r.open()

	var data T
	buffer := (*[]byte)(unsafe.Pointer(into))

	if err := windows.ReadProcessMemory(
		r.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(buffer)),
		uintptr(unsafe.Sizeof(data)),
		nil,
	); err != nil {
		return err
	}

	return nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr Addr, data T) error {
	_ = r.open()

	dtw := (*[]byte)(unsafe.Pointer(&data))

	if err := windows.WriteProcessMemory(
		r.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(dtw)),
		uintptr(unsafe.Sizeof(data)),
		nil,
	); err != nil {
		return err
	}

	return nil
}
