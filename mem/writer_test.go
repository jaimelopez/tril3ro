package mem_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/mem"
)

func TestWriter(t *testing.T) {
	processID := uint32(os.Getpid())

	t.Run("writing uint", func(t *testing.T) {
		var elementToModify uint = 666
		var expected uint = 777

		wtr, _ := mem.NewWriter[uint](mem.WithDefaultHandler(processID))

		err := wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		if err != nil {
			t.Errorf("unexpected error writing into process memory: %s", err)
		}

		if elementToModify != expected {
			t.Errorf("memory incorrect write, expected: %d got %d", expected, elementToModify)
		}
	})

	t.Run("writing string", func(t *testing.T) {
		var elementToModify string = "hi"
		var expected string = "bye"

		wtr, _ := mem.NewWriter[string](mem.WithDefaultHandler(processID))

		err := wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		if err != nil {
			t.Errorf("unexpected error writing into process memory: %s", err)
		}

		if elementToModify != expected {
			t.Errorf("memory incorrect write, expected: %s got %s", expected, elementToModify)
		}
	})

	t.Run("writing struct", func(t *testing.T) {
		type whatever struct {
			ID   int
			Name string
		}

		var elementToModify = whatever{1, "initial name"}
		var expected = whatever{2, "final name"}

		wtr, _ := mem.NewWriter[whatever](mem.WithDefaultHandler(processID))

		err := wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		if err != nil {
			t.Errorf("unexpected error writing into process memory: %s", err)
		}

		if elementToModify != expected {
			t.Errorf("memory incorrect write, expected: %+v got %+v", expected, elementToModify)
		}
	})
}
