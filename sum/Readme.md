# gasoline/sum

The `sum` package is simply a counter of items you pass it. The `New` function
will take any number of arguments. If arguments are passed to `New`, only those
items matching those names will be counted. If no arguments are passed to `New`
every item will be counted.

```go
package main

import "gasoline/sum"

func main() {
  // without arguments
  s := sum.New()

  // a, b, and c are tracked and counted
  s.Insert("a")
  s.Insert("b")
  s.Insert("c")

  // with arguments
  s := sum.New("a", "b")

  // only a and b are tracked and counted. c is ignored
  s.Insert("a")
  s.Insert("b")
  s.Insert("c")
}
```
