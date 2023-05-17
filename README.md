# xjoin: Cartesian Products with Go
A small Go module for performing a cartesian product for a variable number of slices.

This package tries to closely follow Python's [itertools.product()](https://docs.python.org/3/library/itertools.html#itertools.product)

## Example
```go
package main

import (
    "log"
    "github.com/sglmr/xjoin"
)

func main() {
    s1 := []interface{}{"a", "b"}
    s2 := []interface{}{1, 2, 3}

    result := xjoin.Join(s1, s2)
    log.Println(result)
    // [[a 1] [a 2] [a 3] [b 1] [b 2] [b 3]]
}
```

