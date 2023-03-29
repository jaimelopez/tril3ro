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

	// Instantiate a uint32 writer
	r := proc.NewWriter[uint32](p)

	// Calculating the address that we want to read
	addr := mod.Address + 0xD892CC

	// Writting here the calculated address
	err = r.Write(addr, 666)
	if err != nil {
		panic(fmt.Errorf("error writing into addr %x: %s", addr, err.Error()))
	}

	// Successfully written
	fmt.Println("Value written correctly")
}
```
