# Find running processes by name

```go
package proc_example

import (
	"fmt"

	"github.com/jaimelopez/tril3ro/keyboard"
)

func main() {
	// Finding processes with certain name
	procs, err := proc.ProcessByName("csgo_osx64")
	if err != nil {
		panic(fmt.Errorf("error retrieving processes by name: %s", err.Error()))
	}

	// Printing processes.
	// Notice that it is possible to have more than one process with the same name.
	for _, p := range procs {
		fmt.Println(p)
	}
}
```
