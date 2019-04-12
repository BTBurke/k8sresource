# Memcalc

[![Documentation](https://godoc.org/github.com/BTBurke/memcalc?status.svg)](http://godoc.org/github.com/BTBurke/memcalc)

This package contains a few utility functions for converting Kubernetes-style memory definitions between their string and float64 equivalents, and for doing simple math operations.

Example:

```go
package main

import (
    "fmt"
    "log"
    "github.com/BTBurke/memcalc"
)

func main() {
    mem1, err := memcalc.NewFromString("512Mi")
    if err != nil {
        // check error, only supports definitions in Mi or Gi format
        log.Fatalf("unsupported suffix")
    }

    mem2, _ := mem1.Add("512Mi")
    fmt.Printf("As Float: %f\n", mem2.ToFloat64())
    fmt.Printf("As String: %s\n", mem2.ToString())
    //As Float: 1073741824.000000
    //As String: 1024Mi
}
```

See the [godoc](https://godoc.org/github.com/BTBurke/memcalc) for more operations.
