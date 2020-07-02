cpuinfo
=======

Is a simple golang library for get information about CPU on linux or darwin system

Usage
-----

go get github.com/elemc/cpuinfo

```go
package main

import (
    "github.com/elemc/cpuinfo"
    "fmt"
)

func main() {
    info, err := cpuinfo.Get()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Sum for (%s): %2x\n", info.ModelName, info.Sum())
}
	
```