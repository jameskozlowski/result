# Result[T] - A Generic Result Type for Go

Inspired by Rust's `Result<T, E>`, this library provides a generic `Result[T]` type for Go to simplify and clarify error handling through functional-style composition.

## 🚀 Features

- Generic success/error wrapper
- Functional-style chaining: `Map`, `AndThen`, `OrElse`
- Safe unwrapping: `UnwrapOr`, `UnwrapOrElse`, `UnwrapOrDefault`
- Panic-based unwrapping for dev/debug use: `Unwrap`, `Expect`
- Works with any Go type

---

## ✨ Quick Start

```go
func Test(x, y int) Result[int] {
	if x > y {
		return Ok(x)
	}
	return Err[int]("x is not greater than y")
}

func main() {
    res := Test(5, 3).Map(func(v int) int {
	    return v * 2
    })

    if res.IsOk() {
	    fmt.Println("Success:", res.Unwrap())
    } else {
	    fmt.Println("Error:", res.UnwrapErr())
    }

    result2 := Test(2, 3).OrElse(func(err error) Result[int] {
		return Test(5, 3)
	})
	if t.IsOk() {
		println("Result is ok:", t.Unwrap())
	} else {
		println("Error:", t.UnwrapErr().Error())
	}
}
```
