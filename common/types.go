package common

import "strconv"

// ProcessID represents a process identifier
type ProcessID = uint32

// Addr represents a memory address
type Addr = uintptr

// AddrFromString converts a string into an Addr
// Notice that it will not return any error so if anything went wrong, Addr will be 0.
func AddrFromString(addr string) Addr {
	a, _ := strconv.ParseUint(addr, 16, 0)
	return Addr(a)
}
