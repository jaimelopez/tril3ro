package proc_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jaimelopez/tril3ro/proc"
)

func TestProcessByID(t *testing.T) {
	pid := proc.ProcessID(os.Getpid())

	name, err := os.Executable()
	if err != nil {
		t.Errorf("unexpected error retrieving binary name: %s", err)
	}

	name = filepath.Base(name)

	process, err := proc.ProcessByID(pid)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if process.ID != pid {
		t.Errorf("current process id: %d, got: %d", pid, process.ID)
	}

	if process.Name != name {
		t.Errorf("current process name: %s, got: %s", name, process.Name)
	}
}

func TestProcessByName(t *testing.T) {
	pid := proc.ProcessID(os.Getpid())

	name, err := os.Executable()
	if err != nil {
		t.Errorf("unexpected error retrieving binary name: %s", err)
	}

	name = filepath.Base(name)

	processes, err := proc.ProcessByName(name)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(processes) != 1 {
		t.Fatalf("expected to retrieve just 1 process info process, got: %d", len(processes))
	}

	if processes[0].ID != pid {
		t.Errorf("current process id: %d, got: %d", pid, processes[0].ID)
	}

	if processes[0].Name != filepath.Base(name) {
		t.Errorf("current process name: %s, got: %s", name, processes[0].Name)
	}
}
