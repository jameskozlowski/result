package result

import (
	"testing"
)

func TestOrElse(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test OrElse with a successful Result
	result := success.OrElse(func(err error) Result[int] {
		return Ok(0)
	})
	if result.Value != 42 {
		t.Errorf("Expected 42, got %v", result.Value)
	}

	// Test OrElse with an error Result
	result = failure.OrElse(func(err error) Result[int] {
		return Ok(99)
	})
	if result.Value != 99 {
		t.Errorf("Expected 99, got %v", result.Value)
	}
}

func TestExpect(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test Expect with a successful Result
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	if success.Expect("should not panic") != 42 {
		t.Errorf("Expected 42, got %v", success.Value)
	}

	// Test Expect with an error Result
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	_ = failure.Expect("this should panic")
}

func TestExpectErr(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test ExpectErr with an error Result
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	if failure.ExpectErr("should not panic").Error() != "an error occurred" {
		t.Errorf("Expected 'an error occurred', got %v", failure.Err)
	}

	// Test ExpectErr with a successful Result
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	_ = success.ExpectErr("this should panic")
}

func TestUnwrapErr(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test UnwrapErr with an error Result
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Unexpected panic: %v", r)
		}
	}()
	if failure.UnwrapErr().Error() != "an error occurred" {
		t.Errorf("Expected 'an error occurred', got %v", failure.Err)
	}

	// Test UnwrapErr with a successful Result
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	_ = success.UnwrapErr()
}
func TestMap(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test Map with a successful Result
	result := success.Map(func(x int) int {
		return x * 2
	})
	if result.Value != 84 {
		t.Errorf("Expected 84, got %v", result.Value)
	}

	// Test Map with an error Result
	result = failure.Map(func(x int) int {
		return x * 2
	})
	if result.Err.Error() != "an error occurred" {
		t.Errorf("Expected 'an error occurred', got %v", result.Err)
	}
}
func TestMapErr(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test MapErr with a successful Result
	result := success.MapErr(func(err error) error {
		return nil
	})
	if result.Value != 42 {
		t.Errorf("Expected 42, got %v", result.Value)
	}

	// Test MapErr with an error Result
	result = failure.MapErr(func(err error) error {
		return nil
	})
	if result.Err != nil {
		t.Errorf("Expected nil, got %v", result.Err)
	}
}

func TestUnwrap(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test Unwrap with a successful Result
	if success.Unwrap() != 42 {
		t.Errorf("Expected 42, got %v", success.Value)
	}

	// Test Unwrap with an error Result
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but none occurred")
		}
	}()
	_ = failure.Unwrap()
}

func TestUnwrapOr(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test UnwrapOr with a successful Result
	if success.UnwrapOr(0) != 42 {
		t.Errorf("Expected 42, got %v", success.Value)
	}

	// Test UnwrapOr with an error Result
	if failure.UnwrapOr(99) != 99 {
		t.Errorf("Expected 99, got %v", failure.Value)
	}
}

func TestUnwrapOrElse(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test UnwrapOrElse with a successful Result
	if success.UnwrapOrElse(func() int {
		return 0
	}) != 42 {
		t.Errorf("Expected 42, got %v", success.Value)
	}

	// Test UnwrapOrElse with an error Result
	if failure.UnwrapOrElse(func() int {
		return 99
	}) != 99 {
		t.Errorf("Expected 99, got %v", failure.Value)
	}
}

func TestUnwrapOrDefault(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test UnwrapOrDefault with a successful Result
	if success.UnwrapOrDefault() != 42 {
		t.Errorf("Expected 42, got %v", success.Value)
	}

	// Test UnwrapOrDefault with an error Result
	if failure.UnwrapOrDefault() != 0 {
		t.Errorf("Expected 0, got %v", failure.Value)
	}
}

func TestIsOk(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test IsOk with a successful Result
	if !success.IsOk() {
		t.Errorf("Expected true, got false")
	}

	// Test IsOk with an error Result
	if failure.IsOk() {
		t.Errorf("Expected false, got true")
	}
}

func TestIsErr(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test IsErr with a successful Result
	if success.IsErr() {
		t.Errorf("Expected false, got true")
	}

	// Test IsErr with an error Result
	if !failure.IsErr() {
		t.Errorf("Expected true, got false")
	}
}

func TestAndThen(t *testing.T) {
	success := Ok(42)
	failure := Err[int]("an error occurred")

	// Test AndThen with a successful Result
	result := success.AndThen(func(x int) Result[int] {
		return Ok(x * 2)
	})
	if result.Value != 84 {
		t.Errorf("Expected 84, got %v", result.Value)
	}

	// Test AndThen with an error Result
	result = failure.AndThen(func(x int) Result[int] {
		return Ok(x * 2)
	})
	if result.Err.Error() != "an error occurred" {
		t.Errorf("Expected 'an error occurred', got %v", result.Err)
	}
}
