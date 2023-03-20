package proc

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jaimelopez/tril3ro/file"
)

func allProcessesIDs() ([]ProcessID, error) {
	d, err := os.Open("/proc")
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
	filename := fmt.Sprintf("/proc/%d/status", id)

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
