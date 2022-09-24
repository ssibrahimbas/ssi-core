## Ssi Core / Helper

Contains helper functions to write faster and less code.

### CheckErr

Checks the error and exits the program if it is not nil.

type:

```go
func CheckErr(err error)
```

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/helper"
    "fmt"
)

func main() {
    err := fmt.Errorf("error")
    helper.CheckErr(err) // exits the program
}
```