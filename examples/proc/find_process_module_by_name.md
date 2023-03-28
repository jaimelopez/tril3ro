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

	fmt.Println(mod)
}
```
