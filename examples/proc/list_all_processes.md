# List all running processes

```go
package proc_example

import (
	"fmt"

	"github.com/jaimelopez/tril3ro/keyboard"
)

func main() {
    // Listing all running processes
	procs, err := proc.AllProcesses()
	if err != nil {
		panic(fmt.Errorf("error listing all running processes: %s", err.Error()))
	}

    // Printing processes into stdout
	for _, p := range procs {
		fmt.Println(p)
	}
}
```
