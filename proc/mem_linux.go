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

//https://github.com/Andoryuuta/kiwi/blob/214591e6213d10043f39bdbe5eaa84580c4b6679/process_windows.go

// // The platform specific read function.
// func (p *Process) read(addr uintptr, ptr interface{}) error {
// 	// Reflection magic!
// 	v := reflect.ValueOf(ptr)
// 	dataAddr := getDataAddr(v)
// 	dataSize := getDataSize(v)

// 	// Open the file mapped process memory.
// 	mem, err := os.Open(fmt.Sprintf("/proc/%d/mem", p.PID))
// 	defer mem.Close()
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("Error opening /proc/%d/mem. Are you root?", p.PID))
// 	}

// 	// Create a buffer and read data into it.
// 	dataBuf := make([]byte, dataSize)
// 	n, err := mem.ReadAt(dataBuf, int64(addr))
// 	if n != int(dataSize) {
// 		return errors.New(fmt.Sprintf("Tried to read %d bytes, actually read %d bytes\n", dataSize, n))
// 	} else if err != nil {
// 		return err
// 	}

// 	// Unsafely cast to []byte to copy data into.
// 	buf := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
// 		Data: dataAddr,
// 		Len:  int(dataSize),
// 		Cap:  int(dataSize),
// 	}))
// 	copy(*buf, dataBuf)
// 	return nil
// }

// // The platform specific write function.
// func (p *Process) write(addr uintptr, ptr interface{}) error {
// 	// Reflection magic!
// 	v := reflect.ValueOf(ptr)
// 	dataAddr := getDataAddr(v)
// 	dataSize := getDataSize(v)

// 	// Open the file mapped process memory.
// 	mem, err := os.OpenFile(fmt.Sprintf("/proc/%d/mem", p.PID), os.O_WRONLY, 0666)
// 	defer mem.Close()
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("Error opening /proc/%d/mem. Are you root?", p.PID))
// 	}

// 	// Unsafe cast to []byte to copy data from.
// 	buf := (*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
// 		Data: dataAddr,
// 		Len:  int(dataSize),
// 		Cap:  int(dataSize),
// 	}))

// 	// Write the data from buf into memory.
// 	n, err := mem.WriteAt(*buf, int64(addr))
// 	if n != int(dataSize) {
// 		return errors.New((fmt.Sprintf("Tried to write %d bytes, actually wrote %d bytes\n", dataSize, n)))
// 	} else if err != nil {
// 		return err
// 	}
// 	return nil
