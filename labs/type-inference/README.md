## Type inference
In the examples so far we've avoided talking about types, this is because Go supports type
inference.

Type inference lets you omit the type of a variable during declaration.

For example:

```go
var i = 10
```
Go sees that `i` is being declared and initalised with the value 10, so the compiler infers the
type of `i` is an `int`.

However if we did

```go
var name = "ten"
```
Go sees that `name` is being initalised with the string "ten", so the compiler infers the type of `name` is a `string`.

## Explicit type declaration
Sometimes you will want to tell Go to use a specific type.

You do this when you declare a variable:

```go
var i int = 10
var name string = "ten"
```

## Types in Go
Go is a strongly typed language, like Java, C, C++, and Python. Go has nine kinds of types,
they are:

- strings: `string`.
- signed integers: `int8`, `int16`, `int32`, `int64`.
- unsigned integers: `uint8`, `uint16`, `uint32`, `uint64`.
- aliases: `byte`, `rune`, `int`, `uint`.
- booleans: `bool`.
- IEEE floating point: `float32`, `float64`.
- Complex types: `complex64`, `complex128`.
- Compound types: `array`, `slice`, `map`, `struct`.
- Pointer types: `*int`, `*bytes.Buffer`.
 
## String types
Most common data type in Go programs, can be concatenated with the `+` operator

```go
var salutation = "Hello"
var name = "Go"
var greeting = salutation + " " + name
println(greeting)
```
Note: In Go an empty string is "", not `null` or `nil`.

## Integer types
Second most common in Go.

Integer types come in two types; signed and unsigned.
Integer types also come in several sizes, represented by the number of bits they represent;
- Signed integers: `int8`, `int16`, `int32`, `int64`.
- Unsigned integers: `uint8`, `uint16`, `uint32`, `uint64`.

Go has two integer types

- `int`, alias for `int32` or `int64` depending on platform.
- `uint`, alias for `uint32` or `uint64` depending on platform.
whose size depends on the platform you used to build your Go program.

Why does Go support so many kinds of integer types?

<details>
    <summary>Not sure?</summary>
Different sized integer types can accommodate different ranges of numbers. 
Create a go program and paste the below code, compile and run, see what happens:

```go
var x int8 = 400
var y uint = -7
println(x, y)
```
</details>
<br>
