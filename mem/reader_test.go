package mem_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/mem"
)

func TestReader(t *testing.T) {
	processID := uint32(os.Getpid())

	t.Run("reading uint", func(t *testing.T) {
		var elementToRead uint64 = 666
		var elementRetrieved uint64

		reader, _ := mem.NewReader[uint64](mem.WithDefaultHandler(processID))

		err := reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		if err != nil {
			t.Errorf("unexpected error reading process memory: %s", err)
		}

		if elementRetrieved != elementToRead {
			t.Errorf("error reading memory, expected: %d got %d", elementToRead, elementRetrieved)
		}
	})

	t.Run("reading string", func(t *testing.T) {
		var elementToRead string = "this is it"
		var elementRetrieved string

		reader, _ := mem.NewReader[string](mem.WithDefaultHandler(processID))

		err := reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		if err != nil {
			t.Errorf("unexpected error reading process memory: %s", err)
		}

		if elementRetrieved != elementToRead {
			t.Errorf("error reading memory, expected: %s got %s", elementToRead, elementRetrieved)
		}
	})

	t.Run("reading struct", func(t *testing.T) {
		type whatever struct {
			ID   int
			Name string
		}

		elementToRead := whatever{1, "say my name"}
		var elementRetrieved whatever

		reader, _ := mem.NewReader[whatever](mem.WithDefaultHandler(processID))

		err := reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		if err != nil {
			t.Errorf("unexpected error reading process memory: %s", err)
		}

		if elementRetrieved != elementToRead {
			t.Errorf("error reading memory, expected: %+v got %+v", elementToRead, elementRetrieved)
		}
	})
}
