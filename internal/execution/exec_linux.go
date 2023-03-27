package execution

import "syscall"

// IsRoot checks if the process is running under root privileges
func IsRoot() bool {
	return syscall.Getuid() == 0 && syscall.Geteuid() == 0
}
