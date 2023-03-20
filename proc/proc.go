package proc

// ProcessID represents a process identifier
type ProcessID = uint32

// Process definition
type Process struct {
	ID       ProcessID
	ParentID ProcessID
	Name     string
}

// ByID retrieves a process that matches the specified ID
func ByID(id ProcessID) (*Process, error) {
	return process(id)
}

// ByName retrieves a list of process that matches the specified name
func ByName(name string) ([]*Process, error) {
	processes, err := allProcesses()
	if err != nil {
		return nil, err
	}

	matches := []*Process{}

	for _, process := range processes {
		if process.Name == name {
			matches = append(matches, process)
		}
	}

	return matches, nil
}

// All the running processes
func All() ([]*Process, error) {
	return allProcesses()
}

// AllIDs retrieves all the running processes IDs
func AllIDs() ([]ProcessID, error) {
	return allProcessesIDs()
}
