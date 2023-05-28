## Methods

So far we have only been writing functions. Methods are very similar to functions but they are called by invoking them on an instance of a particular type. You can call functions wherever you like, such as `Area(rectangle)` but you can only call methods on instances of type.

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

```go
package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
  return 2*r.Width + 2*r.Height
}

func main() {
  rect := Rectangle{10.0, 15.0}
  fmt.Println(rect.Area())

  // Go automatically handles conversion between values and pointers for method calls. You may
  // want to use a pointer receiver type to avoid copying on method calls or to allow the 
  // method to mutate the receiving struct.

  rp := &rect
  fmt.Println("area: ", rp.Area())
  fmt.Println("perim:", rp.Perimeter())
}
```

[How to decide on choosing a value or a pointer receiver on methods](https://github.com/golang/go/wiki/CodeReviewComments#receiver-type)