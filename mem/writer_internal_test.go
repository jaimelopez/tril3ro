package mem

import (
	"testing"
)

func TestNewWriter(t *testing.T) {
	h := &handler{}

	wrt, err := NewWriter[int](WithtHandler(h))
	if err != nil {
		t.Errorf("unexpected error instantiating new writer: %+v", err)
	}

	if wrt.handler != h {
		t.Errorf("unexpected handler")
	}
}
