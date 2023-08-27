package proc

import (
	"fmt"
	"strings"

	"github.com/jaimelopez/tril3ro/file"
)

// AllModules retrieves all dynamic modules for the process
func (proc *Process) AllModules() ([]*Module, error) {
	scanner, err := file.NewLineScanner(fmt.Sprintf(procMapsLocation, proc.ID))
	if err != nil {
		return nil, err
	}

	defer scanner.Close()

	mods := []*Module{}

	for scanner.Scan() {
		lib := struct {
			StartAddr   string `format:"([[:xdigit:]]*)-.* "`
			EndAddr     string `format:"[[:xdigit:]]*-([[:xdigit:]]*) "`
			Permissions string `format:"[[:xdigit:]]*-[[:xdigit:]]* (.{4}) "`
			Offset      int64  `format:"[[:xdigit:]]*-[[:xdigit:]]* .{4} ([[:xdigit:]]*) "`
			Name        string `format:"/.*/(.*\\.so)(\\..*)?$"`
			Path        string `format:"/.*$"`
		}{}

		scanner.Into(&lib)

		if lib.Name == "" || !strings.Contains(lib.Permissions, "x") {
			continue
		}

		mods = append(mods, &Module{
			Process: proc,
			Address: AddrFromString(lib.StartAddr),
			Size:    uint32(AddrFromString(lib.EndAddr) - AddrFromString(lib.StartAddr)),
			Name:    lib.Name,
			Path:    lib.Path,
		})
	}

	return mods, nil
}
