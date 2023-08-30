package mem

import (
	"os"
	"testing"

	"github.com/jaimelopez/tril3ro/common"
)

func TestNewManager(t *testing.T) {
	handler := &handler{}
	expected := manager{handler}

	withTestingHandler := func(m *manager) error {
		m.handler = handler
		return nil
	}

	m, err := newManager(withTestingHandler)
	if err != nil {
		t.Errorf("unexpected error instantiating new manager: %+v", err)
	}

	if m.handler != handler {
		t.Errorf("unexpected handler")
	}

	if *m != expected {
		t.Errorf("unexpected manager")
	}
}

func TestWithDefaultHandler(t *testing.T) {
	processID := common.ProcessID(os.Getpid())
	mgr := &manager{}
	expected, _ := NewHandler(processID)

	err := WithDefaultHandler(processID)(mgr)
	if err != nil {
		t.Errorf("unexpected error executing option: %+v", err)
	}

	if mgr.handler.processID != expected.processID {
		t.Errorf("unexpected handler")
	}
}

func TestWithtHandler(t *testing.T) {
	processID := common.ProcessID(os.Getpid())
	mgr := &manager{}
	h, _ := NewHandler(processID)

	err := WithtHandler(h)(mgr)
	if err != nil {
		t.Errorf("unexpected error executing option: %+v", err)
	}

	if mgr.handler != h {
		t.Errorf("unexpected handler")
	}
}
