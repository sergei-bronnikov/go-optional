// Package optional provides a type-safe implementation of the Optional pattern
// using Go generics. It offers a container object which may or may not contain
// a value, helping to avoid null pointer dereferences and making the presence
// or absence of values explicit.
//
// Example usage:
//
//	opt := optional.Of("hello")
//	if opt.IsPresent() {
//	    value, _ := opt.Get()
//	    fmt.Println(value)
//	}
package optional

import (
	"fmt"
	"reflect"
)

// Optional is a container object which may or may not contain a value.
// If a value is present, IsPresent() returns true. If no value is present,
// IsEmpty() returns true.
type Optional[T any] struct {
	value *T
}

// Of returns an Optional containing the given value.
//
// Example:
//
//	opt := optional.Of("hello")
func Of[T any](val T) Optional[T] {
	return Optional[T]{value: &val}
}

// OfNullable returns an Optional containing the value pointed to by val if val is not nil,
// otherwise returns an empty Optional.
//
// Example:
//
//	var ptr *string = nil
//	opt := optional.OfNullable(ptr) // empty Optional
//
//	str := "hello"
//	opt = optional.OfNullable(&str) // Optional with "hello"
func OfNullable[T any](val *T) Optional[T] {
	if val == nil {
		return Optional[T]{value: nil}
	}
	return Optional[T]{value: val}
}

// Empty returns an empty Optional instance.
//
// Example:
//
//	opt := optional.Empty[string]()
func Empty[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

// IsPresent returns true if a value is present, otherwise false.
//
// Example:
//
//	opt := optional.Of(42)
//	if opt.IsPresent() {
//	    fmt.Println("Value exists")
//	}
func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

// IsEmpty returns true if no value is present, otherwise false.
//
// Example:
//
//	opt := optional.Empty[string]()
//	if opt.IsEmpty() {
//	    fmt.Println("No value")
//	}
func (o *Optional[T]) IsEmpty() bool {
	return o.value == nil
}

// Get returns the value if present, along with a boolean indicating whether
// the value was present. If no value is present, returns the zero value for type T
// and false.
//
// Example:
//
//	opt := optional.Of("hello")
//	value, ok := opt.Get()
//	if ok {
//	    fmt.Println(value)
//	}
func (o *Optional[T]) Get() (T, bool) {
	if o.value != nil {
		return *o.value, true
	}
	var zero T
	return zero, false
}

// OrElse returns the value if present, otherwise returns the provided default value.
//
// Example:
//
//	opt := optional.Empty[string]()
//	value := opt.OrElse("default") // returns "default"
//
//	opt = optional.Of("hello")
//	value = opt.OrElse("default") // returns "hello"
func (o *Optional[T]) OrElse(other T) T {
	if o.value != nil {
		return *o.value
	}
	return other
}

// Equals compares this Optional with another Optional for equality.
// Two Optionals are considered equal if both are empty, or if both contain
// values that are deeply equal.
//
// Example:
//
//	opt1 := optional.Of(42)
//	opt2 := optional.Of(42)
//	opt1.Equals(opt2) // returns true
func (o *Optional[T]) Equals(other Optional[T]) bool {
	if o.IsPresent() && other.IsPresent() {
		return reflect.DeepEqual(*o.value, *other.value)
	}
	return o.IsEmpty() && other.IsEmpty()
}

// String returns a string representation of the Optional.
// If a value is present, returns "Optional[value]".
// If no value is present, returns "Optional.empty".
//
// Example:
//
//	opt := optional.Of("hello")
//	fmt.Println(opt.String()) // Output: Optional[hello]
func (o *Optional[T]) String() string {
	if o.value != nil {
		return fmt.Sprintf("Optional[%v]", *o.value)
	}
	return "Optional.empty"
}
