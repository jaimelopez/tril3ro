# Keyboard virtual emulate typing

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
	err = virtual.Type("Hello")
	if err != nil {
		panic(fmt.Errorf("error emulating typing using a virtual keyboard: %s", err.Error()))
	}
}
```
