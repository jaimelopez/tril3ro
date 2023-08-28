package proc

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa

 #include <libproc.h>
 #include <mach-o/dyld_images.h>
 #include <mach/mach_traps.h>
 #include <mach/mach_init.h>
*/
import "C"

import (
	"unsafe"
)

const (
	procListPidsMaxSize = 99999
)

// AllProcessesIDs retrieves all the running processes IDs
func AllProcessesIDs() ([]ProcessID, error) {
	bff := make([]ProcessID, procListPidsMaxSize)
	n, err := C.proc_listallpids(unsafe.Pointer(&bff[0]), C.int(len(bff)))

	return bff[:n], err
}

// AllProcesses the running processes
func AllProcesses() ([]*Process, error) {
	pids, err := AllProcessesIDs()
	if err != nil {
		return nil, err
	}

	processes := []*Process{}

	for _, pid := range pids {
		process, err := ProcessByID(pid)
		if err != nil {
			continue
		}

		processes = append(processes, process)
	}

	return processes, nil
}

// ProcessByID retrieves a process that matches the specified ID
func ProcessByID(id ProcessID) (*Process, error) {
	info := &C.struct_proc_taskallinfo{}

	res := C.proc_pidinfo(C.int(id), C.PROC_PIDTASKALLINFO, 0, unsafe.Pointer(info), C.PROC_PIDTASKALLINFO_SIZE)
	if res != C.PROC_PIDTASKALLINFO_SIZE {
		return nil, ErrProcessNotFound
	}

	return &Process{
		ID:       id,
		ParentID: ProcessID(info.pbsd.pbi_ppid),
		Name:     C.GoString(&info.pbsd.pbi_name[0]),
	}, nil
}
