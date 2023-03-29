# Find process module by name

```go
package proc_example

import (
	"fmt"
	"errors"

	"github.com/jaimelopez/tril3ro/keyboard"
)

func main() {
	// Finding processes with a certain name
	procs, err := proc.ProcessByName("csgo_osx64")
	if err != nil {
		panic(fmt.Errorf("error retrieving processes by name: %s", err.Error()))
	}

	if len(procs) == 0 {
		panic(errors.New("no processes found under that name"))
	}

	// For the sake of the example, we just grab the first one
	p := procs[0]

	// Let´s try to find a certain module
	mod, err := p.Module("client.dylib")
	if err != nil {
		panic(fmt.Errorf("error retrieving module: %s", err.Error()))
	}

	// Instantiate a boolean reader
	r := proc.NewReader[bool](p)

	// Calculating the address that we want to read
	addr := mod.Address + 0xD892CC

	// Reading here the calculated address
	value, err := r.Read(addr)
	if err != nil {
		panic(fmt.Errorf("error reader addr %x: %s", addr, err.Error()))
	}

	// Let´s print the boolean value we've just read
	fmt.Println(value)
}
```
