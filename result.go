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

// Ok creates a new Result with a successful value and no error.
func Ok[T any](val T) Result[T] {
	return Result[T]{Value: val, Err: nil}
}

// Err creates a new Result with an error and no value.
func Err[T any](msg string) Result[T] {
	return Result[T]{Err: errors.New(msg)}
}

// Status checks

// IsOk returns true if the Result contains a value and no error.
func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

// IsErr returns true if the Result contains an error.
func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

// Unwrapping

// Unwrap returns the value contained in the Result.
// Panics if the Result contains an error.
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.Err)
	}
	return r.Value
}

// UnwrapOr returns the value contained in the Result, or the provided default value if the Result contains an error.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.Value
}

// UnwrapOrElse returns the value contained in the Result, or the result of the provided function if the Result contains an error.
func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.IsErr() {
		return f()
	}
	return r.Value
}

// UnwrapOrDefault returns the value contained in the Result, or the zero value of the type if the Result contains an error.
func (r Result[T]) UnwrapOrDefault() T {
	var zero T
	if r.IsErr() {
		return zero
	}
	return r.Value
}

// Mapping

// Map applies the provided function to the value contained in the Result, if it is successful.
// Returns a new Result with the transformed value or the original error.
func (r Result[T]) Map(f func(T) T) Result[T] {
	if r.IsErr() {
		return r
	}
	return Ok(f(r.Value))
}

// MapErr applies the provided function to the error contained in the Result, if it contains an error.
// Returns a new Result with the transformed error or the original value.
func (r Result[T]) MapErr(f func(error) error) Result[T] {
	if r.IsOk() {
		return r
	}
	return Result[T]{Err: f(r.Err)}
}

// AndThen applies the provided function to the value contained in the Result, if it is successful.
// Returns the Result returned by the function or the original error.
func (r Result[T]) AndThen(f func(T) Result[T]) Result[T] {
	if r.IsErr() {
		return r
	}
	return f(r.Value)
}

// OrElse applies the provided function to the error contained in the Result, if it contains an error.
// Returns the Result returned by the function or the original value.
func (r Result[T]) OrElse(f func(error) Result[T]) Result[T] {
	if r.IsOk() {
		return r
	}
	return f(r.Err)
}

// Expect-style panics

// Expect returns the value contained in the Result.
// Panics with the provided message if the Result contains an error.
func (r Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(fmt.Sprintf("%s: %v", msg, r.Err))
	}
	return r.Value
}

// ExpectErr returns the error contained in the Result.
// Panics with the provided message if the Result contains a value.
func (r Result[T]) ExpectErr(msg string) error {
	if r.IsOk() {
		panic(msg)
	}
	return r.Err
}

// UnwrapErr returns the error contained in the Result.
// Panics if the Result contains a value.
func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("no error to unwrap")
	}
	return r.Err
}
