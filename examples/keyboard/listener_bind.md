# Keyboard listener bind

```go
package keyboard_example

import (
	"fmt"
	"time"

	"github.com/jaimelopez/tril3ro/keyboard"
)

func main() {
	// Instantiating new listener
	listener, err := keyboard.NewListener()
	if err != nil {
		panic(fmt.Errorf("error instantiating keyboard listener: %s", err.Error()))
	}

	// Binding just certain event
	// If you want to listen for all possible events then better to use listener.Hook() instead
	err = listener.Bind(func() {
		fmt.Println("Key A+B where pressed")
	}, keyboard.NewKeyPressEvent(keyboard.KeyA, keyboard.KeyB))

	// Error produced when hooking
	if err != nil {
		panic(fmt.Errorf("error binding a particular event from keyboard: %s", err.Error()))
	}

	// Let's wait for 10 seconds
	time.Sleep(10 * time.Second)

	// We stop listening
	listener.Stop()
}
```
