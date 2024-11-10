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
