package proc

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa

 #include <libproc.h>
*/
import "C"

import (
	"unsafe"
)

const (
	procListPidsMaxSize = 99999
)

func allProcessesIDs() ([]ProcessID, error) {
	bff := make([]ProcessID, procListPidsMaxSize)
	n, err := C.proc_listallpids(unsafe.Pointer(&bff[0]), C.int(len(bff)))

	return bff[:n], err
}

func allProcesses() ([]*Process, error) {
	pids, err := allProcessesIDs()
	if err != nil {
		return nil, err
	}

	processes := []*Process{}

	for _, pid := range pids {
		process, err := process(pid)
		if err != nil {
			continue
		}

		processes = append(processes, process)
	}

	return processes, nil
}

func process(id ProcessID) (*Process, error) {
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
