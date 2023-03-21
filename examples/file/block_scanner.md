# File block scanner

In this example our file contains two blocks with two properties each so the sample data would look like this:

```text
Name: 			Name of the first block
Description: 	Description of the first block

Name: 			Name of the second block
Description: 	Description of the second block
```

After scanning and parsing the files, the aim is to end up having a list of `struct{Name, Description}`

```go
package file_examples

// Format tag needs to be specified in order to populate the match into the struct property.
// Expected format is a regex with the match group.
// If more than one match are found, then the first one is going to be populated
type BlockExample struct {
	Name string `format:"Name:\\s*(.*)"`
	Description uint `format:"Description:\\s*(.*)"`
}

func blockScanner() {
	r, err := file.NewBlockScanner("test.file")
	if err != nil {
		panic(fmt.Errorf("error instantiating new block scanner: %s", err.Error()))
	}
	defer r.Close()

	blocks := []BlockExamples{}

	for r.Scan() {
		current := BlockExample{}

		err := r.Into(&current)
		if err != nil {
			continue
		}

		blocks = append(blocks, current)
	}

	fmt.Println(blocks)
}
```
