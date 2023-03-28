package proc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaimelopez/tril3ro/file"
)

const (
	procsLocation      = "/proc"
	procStatusLocation = "/proc/%d/status"
	procMapsLocation   = "/proc/%d/maps"
	procMemLocation    = "/proc/%d/mem"
)

// AllProcessesIDs retrieves all the running processes IDs
func AllProcessesIDs() ([]ProcessID, error) {
	d, err := os.Open(procsLocation)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return nil, ErrOperationNotAllowed
	}

	procs := []ProcessID{}

	for _, n := range names {
		pid, err := strconv.ParseUint(n, 10, 32)
		if err != nil {
			continue
		}

		procs = append(procs, ProcessID(pid))
	}

	return procs, nil
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
	filename := fmt.Sprintf(procStatusLocation, id)

	s, err := file.NewBlockScanner(filename)
	if err != nil {
		return nil, ErrProcessNotFound
	}

	defer s.Close()

	ls := struct {
		Name string    `format:"Name:\\s*(.*)"`
		PPid ProcessID `format:"PPid:.\\s*(\\d*)"`
	}{}

	if s.Scan() {
		_ = s.Into(&ls)
	}

	return &Process{
		ID:       id,
		ParentID: ls.PPid,
		Name:     ls.Name,
	}, nil
}
