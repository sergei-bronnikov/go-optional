package optional

import (
	"fmt"
	"reflect"
)

type Optional[T any] struct {
	value *T
}

func Of[T any](val T) Optional[T] {
	return Optional[T]{value: &val}
}

func OfNullable[T any](val *T) Optional[T] {
	if val == nil {
		return Optional[T]{value: nil}
	}
	return Optional[T]{value: val}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

func (o *Optional[T]) IsPresent() bool {
	return o.value != nil
}

func (o *Optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o *Optional[T]) Get() (T, bool) {
	if o.value != nil {
		return *o.value, true
	}
	var zero T
	return zero, false
}

func (o *Optional[T]) OrElse(other T) T {
	if o.value != nil {
		return *o.value
	}
	return other
}

func (o *Optional[T]) Equals(other Optional[T]) bool {
	if o.IsPresent() && other.IsPresent() {
		return reflect.DeepEqual(*o.value, *other.value)
	}
	return o.IsEmpty() && other.IsEmpty()
}

func (o *Optional[T]) String() string {
	if o.value != nil {
		return fmt.Sprintf("Optional[%v]", *o.value)
	}
	return "Optional.empty"
}
