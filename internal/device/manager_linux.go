package device

import (
	"fmt"

	"github.com/jaimelopez/tril3ro/file"
)

const (
	procDevices = "/proc/bus/input/devices"
	devInput    = "/dev/input/%s"

	evKeyboardCapabilities = 120013
)

func FindKeyboardEvents() ([]string, error) {
	s, err := file.NewBlockScanner(procDevices)
	if err != nil {
		return nil, err
	}

	events := []string{}

	for s.Scan() {
		block := struct {
			EV      uint   `format:"B: EV=([[:xdigit:]]*)"`
			Handler string `format:"H: Handlers=.*(event\\d*)"`
		}{}

		s.Into(&block)

		if (block.EV & evKeyboardCapabilities) != evKeyboardCapabilities {
			continue
		}

		if block.Handler != "" {
			events = append(events, fmt.Sprintf(devInput, block.Handler))
		}
	}

	return events, nil
}
