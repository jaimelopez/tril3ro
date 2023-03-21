# Keyboard virtual press and release

```go
package keyboard_example

import (
	"fmt"

	"github.com/jaimelopez/tril3ro/keyboard"
)

func main() {
	// Instantiating new virtual keyboard
	virtual, err := keyboard.NewVirtual()
	if err != nil {
		panic(fmt.Errorf("error instantiating virtual keyboard: %s", err.Error()))
	}

	// We just emulate that we press and released 'HELLO'
	err = virtual.PressAndRelease(keyboard.KeyH, keyboard.KeyE, keyboard.KeyL, keyboard.KeyL, keyboard.KeyO)
	if err != nil {
		panic(fmt.Errorf("error emulating pressing few keys with virtual keyboard: %s", err.Error()))
	}
}
```
