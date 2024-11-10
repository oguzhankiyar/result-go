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
