package proc_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

func BenchmarkRead(b *testing.B) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	b.Run("reading uint", func(b *testing.B) {
		var elementToRead uint64 = 666
		var elementRetrieved uint64

		reader := proc.NewReader[uint64](process)

		for i := 0; i < b.N; i++ {
			_ = reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		}
	})

	b.Run("reading string", func(b *testing.B) {
		var elementToRead string = "this is it"
		var elementRetrieved string

		reader := proc.NewReader[string](process)

		for i := 0; i < b.N; i++ {
			_ = reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		}
	})

	b.Run("reading struct", func(b *testing.B) {
		type whatever struct {
			ID   int
			Name string
		}

		elementToRead := whatever{1, "say my name"}
		var elementRetrieved whatever

		reader := proc.NewReader[whatever](process)

		for i := 0; i < b.N; i++ {
			_ = reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		}
	})
}

func BenchmarkWriter(b *testing.B) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	b.Run("writing uint", func(b *testing.B) {
		var elementToModify uint = 666
		var expected uint = 777

		wtr := proc.NewWriter[uint](process)

		for i := 0; i < b.N; i++ {
			_ = wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		}
	})

	b.Run("writing string", func(b *testing.B) {
		var elementToModify string = "hi"
		var expected string = "bye"

		wtr := proc.NewWriter[string](process)

		for i := 0; i < b.N; i++ {
			_ = wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		}
	})

	b.Run("writing struct", func(b *testing.B) {
		type whatever struct {
			ID   int
			Name string
		}

		var elementToModify = whatever{1, "initial name"}
		var expected = whatever{2, "final name"}

		wtr := proc.NewWriter[whatever](process)

		for i := 0; i < b.N; i++ {
			_ = wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		}
	})
}
