package device

import (
	"fmt"

	"github.com/jaimelopez/tril3ro/file"
)

const (
	devicesLocation     = "/proc/bus/input/devices"
	deviceInputLocation = "/dev/input/%s"

	evKeyboardCapabilities = 120013
)

func FindKeyboardEvents() ([]string, error) {
	scanner, err := file.NewBlockScanner(devicesLocation)
	if err != nil {
		return nil, err
	}

	events := []string{}

	for scanner.Scan() {
		block := struct {
			EV      uint   `format:"B: EV=([[:xdigit:]]*)"`
			Handler string `format:"H: Handlers=.*(event\\d*)"`
		}{}

		scanner.Into(&block)

		if (block.EV & evKeyboardCapabilities) != evKeyboardCapabilities {
			continue
		}

		if block.Handler != "" {
			events = append(events, fmt.Sprintf(deviceInputLocation, block.Handler))
		}
	}

	return events, nil
}
