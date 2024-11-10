# Result

This package provides a `Result` type that encapsulates the outcome of an operation, handling both successful and error results. The `Result` type is generic, enabling its use with any data type. It simplifies error handling and allows you to build functional-style chains for more readable and expressive code.

## Installation

To use this package, add it to your Go project.

```bash
go get github.com/oguzhankiyar/result-go
```

You can import it like below.

```go
import "github.com/oguzhankiyar/result-go"
```

## Usage

### Basic Structure

The `Result` struct has two main fields:
- `val`: Stores the value if the operation is successful.
- `err`: Stores the error if the operation fails.

### Creating Results

You can create `Result` instances using the helper functions:

- `Val(val T) Result[T]`: Creates a successful result containing `val`.
- `Err(err error) Result[T]`: Creates a failed result with the specified error.
- `Wrap(val T, err error) Result[T]`: Wraps a value and an error, creating a `Result`. If `err` is not nil, it creates a failed result.

### Methods

#### Ok
```go
func (res Result[T]) Ok() bool
```
Returns `true` if the result is successful (no error).

#### Val
```go
func (res Result[T]) Val() T
```
Returns the value if the result is successful, or the zero value of `T` if there’s an error.

#### Err
```go
func (res Result[T]) Err() error
```
Returns the error if there is one, or nil if the result is successful.

#### Unwrap
```go
func (res Result[T]) Unwrap() (T, error)
```
Returns a tuple of the value and the error, allowing you to handle both in a single statement.

#### Fallback
```go
func (res Result[T]) Fallback(fallback T) Result[T]
```
Returns the original result if successful, or a result with the `fallback` value if there's an error.

#### Ensure
```go
func (res Result[T]) Ensure(condition func(T) bool, err error) Result[T]
```
Ensures that a condition is met. If not, returns a new failed result with the provided error.

#### Tap
```go
func (res Result[T]) Tap(onSuccess func(T), onError func(error)) Result[T]
```
Calls `onSuccess` if the result is successful, or `onError` if it failed. Returns the original result.

### Functional Helpers

#### Pipe
```go
func Pipe[T any, U any](res Result[T], fn func(T) (U, error)) Result[U]
```
Applies a function to the result’s value if successful. If the function returns an error, the result will be an error; otherwise, the transformed value is returned.

#### Map
```go
func Map[T, U any](res Result[T], fn func(T) U) Result[U]
```
Transforms the value of the result if successful, using a function that cannot fail.

#### Fold
```go
func Fold[T, U any](res Result[T], onSuccess func(T) U, onError func(error) U) U
```
Unwraps the result, applying `onSuccess` if successful, or `onError` if failed, and returns the final value.

## Example
```go
package main

import (
	"errors"
	"fmt"
	
	"github.com/oguzhankiyar/result-go"
)

func main() {
	// Attempt to divide 10 by 2 (successful operation)
	// Output: Result: 5
	res1 := divide(10, 2)
	if res1.Ok() {
		fmt.Println("Result:", res1.Val())
	} else {
		fmt.Println("Error:", res1.Err())
	}

	// Attempt to divide 10 by 0 (error case)
	// Output: Error: division by zero
	res2 := divide(10, 0)
	if res2.Ok() {
		fmt.Println("Result:", res2.Val())
	} else {
		fmt.Println("Error:", res2.Err())
	}
}

func divide(a, b int) result.Result[int] {
	if b == 0 {
		return result.Err[int](errors.New("division by zero"))
	}
	return result.Val(a / b)
}
```

## License
This library is licensed under the [MIT License](LICENSE).