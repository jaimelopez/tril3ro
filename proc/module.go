package proc

// Module represents a dynamic module inside a process
type Module struct {
	ProcessID ProcessID
	Address   Addr
	Size      uint32
	Name      string
	Path      string
}
