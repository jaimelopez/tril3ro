package proc

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	procListPidsMaxSize = 99999
)

func allProcessesIDs() ([]ProcessID, error) {
	pIDs := make([]uint32, procListPidsMaxSize)
	var result uint32

	err := windows.EnumProcesses(pIDs, &result)
	if err != nil {
		return nil, err
	}

	if unsafe.Sizeof(result) <= 0 {
		return nil, ErrUnexpectedResult
	}

	// return slices.Clip(pids), nil

	return pIDs[:uintptr(result)/unsafe.Sizeof(result)], nil
}

func allProcesses() ([]*Process, error) {
	handle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}

	defer windows.CloseHandle(handle)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	err = windows.Process32First(handle, &entry)
	if err != nil {
		return nil, err
	}

	processes := []*Process{}

	for {
		processes = append(processes, &Process{
			ID:       ProcessID(entry.ProcessID),
			ParentID: ProcessID(entry.ParentProcessID),
			Name:     syscall.UTF16ToString(entry.ExeFile[:]),
		})

		err := windows.Process32Next(handle, &entry)
		if err != nil {
			break
		}
	}

	return processes, nil
}

func process(id ProcessID) (*Process, error) {
	processes, err := allProcesses()
	if err != nil {
		return nil, err
	}

	for _, process := range processes {
		if process.ID == id {
			return process, nil
		}
	}

	return nil, ErrProcessNotFound
}
