package mem_test

import (
	"os"
	"testing"
	"unsafe"

	"github.com/jaimelopez/tril3ro/mem"
)

func BenchmarkWriter(b *testing.B) {
	processID := uint32(os.Getpid())

	b.Run("writing uint", func(b *testing.B) {
		var elementToModify uint = 666
		var expected uint = 777

		wtr, _ := mem.NewWriter[uint](mem.WithDefaultHandler(processID))

		for i := 0; i < b.N; i++ {
			_ = wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		}
	})

	b.Run("writing string", func(b *testing.B) {
		var elementToModify string = "hi"
		var expected string = "bye"

		wtr, _ := mem.NewWriter[string](mem.WithDefaultHandler(processID))

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

		wtr, _ := mem.NewWriter[whatever](mem.WithDefaultHandler(processID))

		for i := 0; i < b.N; i++ {
			_ = wtr.Write(uintptr(unsafe.Pointer(&elementToModify)), expected)
		}
	})
}
