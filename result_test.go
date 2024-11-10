package result

import (
	"errors"
	"strconv"
	"testing"
)

func TestResult(t *testing.T) {
	t.Run("Test Val Success", func(t *testing.T) {
		res := Val(42)
		if !res.Ok() {
			t.Errorf("Expected success, got error: %v", res.Err())
		}
		if res.Val() != 42 {
			t.Errorf("Expected value 42, got: %v", res.Val())
		}
	})

	t.Run("Test Err Failure", func(t *testing.T) {
		expectedError := errors.New("something went wrong")
		res := Err[int](expectedError)
		if res.Ok() {
			t.Error("Expected failure, got success")
		}
		if res.Err() == nil {
			t.Error("Expected error to be non-nil")
		}
		if res.Err().Error() != expectedError.Error() {
			t.Errorf("Expected error '%v', got: '%v'", expectedError, res.Err())
		}
	})

	t.Run("Test Fallback on Error", func(t *testing.T) {
		expectedError := errors.New("something went wrong")
		res := Err[int](expectedError)
		fallback := 100
		result := res.Fallback(fallback)

		if res.Ok() {
			t.Error("Expected failure, got success")
		}
		if result.Val() != fallback {
			t.Errorf("Expected fallback value %d, got %d", fallback, result.Val())
		}
	})

	t.Run("Test Ensure Success", func(t *testing.T) {
		res := Val(42)
		condition := func(val int) bool {
			return val > 0
		}
		res = res.Ensure(condition, errors.New("value must be positive"))

		if !res.Ok() {
			t.Errorf("Expected success, got error: %v", res.Err())
		}
	})

	t.Run("Test Ensure Failure", func(t *testing.T) {
		res := Val(42)
		condition := func(val int) bool {
			return val < 0
		}
		expectedError := errors.New("value must be positive")
		res = res.Ensure(condition, expectedError)

		if res.Ok() {
			t.Error("Expected failure, got success")
		}
		if res.Err() == nil || res.Err().Error() != expectedError.Error() {
			t.Errorf("Expected error '%v', got: '%v'", expectedError, res.Err())
		}
	})

	t.Run("Test Tap Success", func(t *testing.T) {
		res := Val(42)
		var tappedValue int
		res.Tap(func(val int) {
			tappedValue = val
		}, func(err error) {
			t.Errorf("Expected success, but got error: %v", err)
		})

		if tappedValue != 42 {
			t.Errorf("Expected tapped value 42, got %d", tappedValue)
		}
	})

	t.Run("Test Tap Failure", func(t *testing.T) {
		expectedError := errors.New("something went wrong")
		res := Err[int](expectedError)
		var tappedError error
		res.Tap(func(val int) {
			t.Errorf("Expected error, but got value %d", val)
		}, func(err error) {
			tappedError = err
		})

		if tappedError == nil || tappedError.Error() != expectedError.Error() {
			t.Errorf("Expected error '%v', got: '%v'", expectedError, tappedError)
		}
	})

	t.Run("Test Pipe Success", func(t *testing.T) {
		res := Val(42)
		pipeResult := Pipe(res, func(val int) (string, error) {
			return "Value: " + strconv.Itoa(val), nil
		})

		if !pipeResult.Ok() {
			t.Errorf("Expected success, got error: %v", pipeResult.Err())
		}
		if pipeResult.Val() != "Value: 42" {
			t.Errorf("Expected piped value 'Value: 42', got: %v", pipeResult.Val())
		}
	})

	t.Run("Test Pipe Failure", func(t *testing.T) {
		expectedError := errors.New("pipe error")
		res := Err[int](nil)
		pipeResult := Pipe(res, func(val int) (string, error) {
			return "", expectedError
		})

		if pipeResult.Ok() {
			t.Error("Expected failure, got success")
		}
		if pipeResult.Err() == nil || pipeResult.Err().Error() != expectedError.Error() {
			t.Errorf("Expected error '%v', got: '%v'", expectedError, pipeResult.Err())
		}
	})

	t.Run("Test Map Success", func(t *testing.T) {
		res := Val(42)
		mapResult := Map(res, func(val int) string {
			return "Value: " + strconv.Itoa(val)
		})

		if !mapResult.Ok() {
			t.Errorf("Expected success, got error: %v", mapResult.Err())
		}
		if mapResult.Val() != "Value: 42" {
			t.Errorf("Expected mapped value 'Value: 42', got: %v", mapResult.Val())
		}
	})

	t.Run("Test Map Failure", func(t *testing.T) {
		expectedError := errors.New("map error")
		res := Err[int](expectedError)
		mapResult := Map(res, func(val int) string {
			return ""
		})

		if mapResult.Ok() {
			t.Error("Expected failure, got success")
		}
		if mapResult.Err() == nil || mapResult.Err().Error() != expectedError.Error() {
			t.Errorf("Expected error '%v', got: '%v'", expectedError, mapResult.Err())
		}
	})

	t.Run("Test Fold Success", func(t *testing.T) {
		res := Val(42)
		result := Fold(res, func(val int) string {
			return "Value: " + strconv.Itoa(val)
		}, func(err error) string {
			return "Error: " + err.Error()
		})

		if result != "Value: 42" {
			t.Errorf("Expected result 'Value: 42', got: %v", result)
		}
	})

	t.Run("Test Fold Failure", func(t *testing.T) {
		expectedError := errors.New("something went wrong")
		res := Err[int](expectedError)
		result := Fold(res, func(val int) string {
			return "Value: " + strconv.Itoa(val)
		}, func(err error) string {
			return "Error: " + err.Error()
		})

		if result != "Error: something went wrong" {
			t.Errorf("Expected result 'Error: something went wrong', got: %v", result)
		}
	})
}
