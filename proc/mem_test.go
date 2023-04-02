package proc_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

func TestReader(t *testing.T) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	t.Run("reading uint", func(t *testing.T) {
		var elementToRead uint64 = 666
		var elementRetrieved uint64

		reader := proc.NewReader[uint64](process)

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

		reader := proc.NewReader[string](process)

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

		reader := proc.NewReader[whatever](process)

		err := reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		if err != nil {
			t.Errorf("unexpected error reading process memory: %s", err)
		}

		if elementRetrieved != elementToRead {
			t.Errorf("error reading memory, expected: %+v got %+v", elementToRead, elementRetrieved)
		}
	})
}

func TestNewReader(t *testing.T) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	rdr := proc.NewReader[int](process)
	expected := proc.Reader[int]{process}

	if *rdr != expected {
		t.Errorf("unexpected reader")
	}
}

func TestWriter(t *testing.T) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	t.Run("writing uint", func(t *testing.T) {
		var elementToModify uint = 666
		var expected uint = 777

		wtr := proc.NewWriter[uint](process)

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

		wtr := proc.NewWriter[string](process)

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

		wtr := proc.NewWriter[whatever](process)

		err := wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		if err != nil {
			t.Errorf("unexpected error writing into process memory: %s", err)
		}

		if elementToModify != expected {
			t.Errorf("memory incorrect write, expected: %+v got %+v", expected, elementToModify)
		}
	})
}

func TestProcWrite(t *testing.T) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	t.Run("writing uint", func(t *testing.T) {
		var elementToModify uint = 666
		var expected uint = 777

		err := process.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
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

		err := process.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
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

		err := process.Write(uintptr(unsafe.Pointer(&elementToModify)), &expected)
		if err != nil {
			t.Errorf("unexpected error writing into process memory: %s", err)
		}

		if elementToModify != expected {
			t.Errorf("memory incorrect write, expected: %+v got %+v", expected, elementToModify)
		}
	})
}

func TestNewWriter(t *testing.T) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	wrt := proc.NewWriter[int](process)
	expected := proc.Writer[int]{process}

	if *wrt != expected {
		t.Errorf("unexpected writer")
	}
}
