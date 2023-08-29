package mem_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/mem"
	"github.com/jaimelopez/tril3ro/proc"
)

func BenchmarkRead(b *testing.B) {
	process, _ := proc.ProcessByID(uint32(os.Getpid()))

	b.Run("reading uint", func(b *testing.B) {
		var elementToRead uint64 = 666
		var elementRetrieved uint64

		reader, _ := mem.NewReader[uint64](mem.WithDefaultHandler(process.ID))

		for i := 0; i < b.N; i++ {
			_ = reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		}
	})

	b.Run("reading string", func(b *testing.B) {
		var elementToRead string = "this is it"
		var elementRetrieved string

		reader, _ := mem.NewReader[string](mem.WithDefaultHandler(process.ID))

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

		reader, _ := mem.NewReader[whatever](mem.WithDefaultHandler(process.ID))

		for i := 0; i < b.N; i++ {
			_ = reader.Read(uintptr(unsafe.Pointer(&elementToRead)), &elementRetrieved)
		}
	})
}
