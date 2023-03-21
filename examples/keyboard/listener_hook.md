# Keyboard listener hook

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

	// Hooking for all events, we can perform some operations and checks inside the func if needed
	// If you want to just hook a particular event (certain key for instance), better to use listener.Bind() instead
	err = listener.Hook(func(event keyboard.Event) {
		fmt.Println("Event type", event.Type, "received from keyboard:", event.Keys)
	})

	// Error produced when hooking
	if err != nil {
		panic(fmt.Errorf("error hooking all events from keyboard: %s", err.Error()))
	}

	// Let's wait for 10 seconds
	time.Sleep(10 * time.Second)

	// We stop listening
	listener.Stop()
}
```
