package mem

import (
	"testing"
)

func TestNewReader(t *testing.T) {
	h := &handler{}

	wrt, err := NewReader[int](WithtHandler(h))
	if err != nil {
		t.Errorf("unexpected error instantiating new reader: %+v", err)
	}

	if wrt.handler != h {
		t.Errorf("unexpected handler")
	}
}
