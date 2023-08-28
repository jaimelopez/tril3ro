package proc

// Module represents a dynamic module inside a process
type Module struct {
	Process *Process
	Address Addr
	Size    uint32
	Name    string
	Path    string
}

// func (m *Module) ReadAll() ([]byte, error) {
// 	rdr := NewReader[[]byte](m.Process)

// 	data := make([]byte, m.Size)

// 	err := rdr.ReadOf(m.Address, &data, uint(m.Size))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
