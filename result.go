package result

import (
	"errors"
	"fmt"
)

type Result[T any] struct {
	Value T
	Err   error
}

// Constructors
func Ok[T any](val T) Result[T] {
	return Result[T]{Value: val, Err: nil}
}

func Err[T any](msg string) Result[T] {
	return Result[T]{Err: errors.New(msg)}
}

// Status checks
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

// Unwrapping
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.Err)
	}
	return r.Value
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.Value
}

func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.IsErr() {
		return f()
	}
	return r.Value
}

func (r Result[T]) UnwrapOrDefault() T {
	var zero T
	if r.IsErr() {
		return zero
	}
	return r.Value
}

// Mapping
func (r Result[T]) Map(f func(T) T) Result[T] {
	if r.IsErr() {
		return r
	}
	return Ok(f(r.Value))
}

func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.IsOk() {
		return r
	}
	return Result[T]{Err: f(r.Err)}
}

func (r Result[T]) AndThen(f func(T) Result[T]) Result[T] {
	if r.IsErr() {
		return r
	}
	return f(r.Value)
}

func (r Result[T]) OrElse(f func(error) Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return f(r.Err)
}

// Expect-style panics
func (r Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(fmt.Sprintf("%s: %v", msg, r.Err))
	}
	return r.Value
}

func (r Result[T]) ExpectErr(msg string) error {
	if r.IsOk() {
		panic(msg)
	}
	return r.Err
}

func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("no error to unwrap")
	}
	return r.Err
}
