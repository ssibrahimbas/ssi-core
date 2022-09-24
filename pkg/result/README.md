## Ssi Core / Result

This package has been developed to ensure that each microservice returns a response using the same result type. Thus, there will be consistency between microservices.

### Types

#### result

```go
type Result struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type DataResult struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
```

Result is a struct that includes the result of the operation.

### Functions

```go
func Success(m string, c int) *Result
func Error(m string, c int) *Result
func SuccessData(m string, d any, c int) *DataResult
func ErrorData(m string, d any, c int) *DataResult
```

### Create a Result

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/result"
    "fmt"
)

func main() {
    r := result.Success("success", 200)
    fmt.Println(r)
}
```

### Create a Data Result

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/result"
    "fmt"
)

func main() {
    r := result.SuccessData("success", map[string]string{
        "name": "Ssi Core",
    }, 200)
    fmt.Println(r)
}
```

