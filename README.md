# logSys
A simple to use logo system, minimalist but with features for debugging and differentiation of messages


## Example

[![Example](examples/example.png)](examples/example.go)


```go
package main

import (
    "fmt"
    "github.com/nuveo/logSys"
)

func main() {
    log.Debugln("Debug message")

    log.DebugMode = false
    log.Debugln("Debug message that will be hidden")

    log.Println("Info message")
    log.Warningln("Warning message")
    log.Errorln("Error message")
    log.Fatal("Fatal error message")
    fmt.Println("I will never be printed because of Fatal()")
}
```