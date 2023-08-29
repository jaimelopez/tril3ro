package proc

import (
	"github.com/jaimelopez/tril3ro/common"
	"github.com/jaimelopez/tril3ro/mem"
)

// Module represents a dynamic module inside a process
type Module struct {
	ProcessID common.ProcessID
	Address   common.Addr
	Size      uint32
	Name      string
	Path      string
}

func (m *Module) ReadAll() ([]byte, error) {
	rdr, err := mem.NewReaderForProc[[]byte](m.ProcessID)
	if err != nil {
		return nil, err
	}

	data := make([]byte, m.Size)

	err = rdr.ReadOf(m.Address, &data, uint(m.Size))
	if err != nil {
		return nil, err
	}

	return data, nil
}
