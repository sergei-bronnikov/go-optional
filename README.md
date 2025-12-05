# go-optional

A Go implementation of the Optional pattern using generics, inspired by Java's Optional class. This library provides a type-safe way to handle optional values and avoid null pointer dereferences.

## Features

- Type-safe optional values using Go generics
- Familiar API for developers coming from Java
- Zero dependencies
- Simple and intuitive interface

## Installation

```bash
go get github.com/sergei-bronnikov/go-optional
```

## Usage

### Creating Optional Values

```go
import "github.com/sergei-bronnikov/go-optional"

// Create an Optional with a value
opt := optional.Of("hello")

// Create an Optional from a pointer (can be nil)
var ptr *string
opt := optional.OfNullable(ptr)

// Create an empty Optional
opt := optional.Empty[string]()
```

### Checking for Values

```go
opt := optional.Of(42)

if opt.IsPresent() {
    fmt.Println("Value is present")
}

if opt.IsEmpty() {
    fmt.Println("Value is empty")
}
```

### Getting Values

```go
opt := optional.Of("hello")

// Get with boolean indicator
value, ok := opt.Get()
if ok {
    fmt.Println(value)
}

// Get with default value
value := opt.OrElse("default")
```

### Comparing Optionals

```go
opt1 := optional.Of(42)
opt2 := optional.Of(42)
opt3 := optional.Empty[int]()

opt1.Equals(opt2) // true
opt1.Equals(opt3) // false
```

### String Representation

```go
opt := optional.Of("hello")
fmt.Println(opt.String()) // Output: Optional[hello]

empty := optional.Empty[string]()
fmt.Println(empty.String()) // Output: Optional.empty
```

## API Reference

### Types

- `Optional[T any]` - A container object which may or may not contain a value

### Functions

- `Of[T any](val T) Optional[T]` - Creates an Optional with the given non-nil value
- `OfNullable[T any](val *T) Optional[T]` - Creates an Optional from a pointer, which may be nil
- `Empty[T any]() Optional[T]` - Returns an empty Optional

### Methods

- `IsPresent() bool` - Returns true if a value is present
- `IsEmpty() bool` - Returns true if no value is present
- `Get() (T, bool)` - Returns the value if present, and a boolean indicating presence
- `OrElse(other T) T` - Returns the value if present, otherwise returns the provided default
- `Equals(other Optional[T]) bool` - Compares two Optionals for equality
- `String() string` - Returns a string representation of the Optional

## License

See LICENSE file for details.
