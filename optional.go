package optional

type Maybe[T any] interface {
	Get() T
	OrElse(T) T
	OrElseGet(func() T) T

	IsPresent() bool
	IfPresent(func(T))
}

type Optional[T any] struct {
	Maybe[T]
	value *T
}

func (opt Optional[T]) Get() T {
	return *opt.value
}

func (opt Optional[T]) OrElse(other T) T {
	if opt.IsPresent() {
		return *opt.value
	}
	return other
}

func (opt Optional[T]) OrElseGet(supplier func() T) T {
	if opt.IsPresent() {
		return *opt.value
	}
	return supplier()
}

func (opt Optional[T]) IsPresent() bool {
	return opt.value != nil
}

func (opt Optional[T]) IfPresent(consumer func(T)) {
	if opt.IsPresent() {
		consumer(*opt.value)
	}
}

func Of[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}
