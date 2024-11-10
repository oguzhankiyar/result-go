package result

type Result[T any] struct {
	val T
	err error
}

func (res Result[T]) Ok() bool {
	return res.err == nil
}

func (res Result[T]) Val() T {
	if res.err != nil {
		return *new(T)
	}
	return res.val
}

func (res Result[T]) Err() error {
	return res.err
}

func (res Result[T]) Unwrap() (T, error) {
	return res.val, res.err
}

func (res Result[T]) Fallback(fallback T) Result[T] {
	if res.Ok() {
		return res
	}
	return Val(fallback)
}

func (res Result[T]) Ensure(condition func(T) bool, err error) Result[T] {
	if !res.Ok() || condition(res.Val()) {
		return res
	}
	return Err[T](err)
}

func (res Result[T]) Tap(onSuccess func(T), onError func(error)) Result[T] {
	if res.Ok() {
		onSuccess(res.val)
	} else {
		onError(res.Err())
	}
	return res
}

func Val[T any](val T) Result[T] {
	return Result[T]{
		val: val,
		err: nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		val: *new(T),
		err: err,
	}
}

func Wrap[T any](val T, err error) Result[T] {
	return Result[T]{
		val: val,
		err: err,
	}
}

func Pipe[T any, U any](res Result[T], fn func(T) (U, error)) Result[U] {
	if res.Ok() {
		newVal, err := fn(res.val)
		if err != nil {
			return Err[U](err)
		}
		return Val[U](newVal)
	}
	return Err[U](res.Err())
}

func Map[T, U any](res Result[T], fn func(T) U) Result[U] {
	if res.Ok() {
		return Val(fn(res.val))
	}
	return Err[U](res.Err())
}

func Fold[T, U any](res Result[T], onSuccess func(T) U, onError func(error) U) U {
	if res.Ok() {
		return onSuccess(res.val)
	}
	return onError(res.Err())
}
