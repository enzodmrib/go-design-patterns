package iterator

type Iterator[T any] struct {
	index             int
	list              []T
	wasIteratorCalled bool
}

func NewIterator[T any](list []T) *Iterator[T] {
	return &Iterator[T]{
		index:             0,
		list:              list,
		wasIteratorCalled: false,
	}
}

func (i *Iterator[T]) Get() T {
	return i.list[i.index]
}

func (i *Iterator[T]) Next() bool {
	if i.index < len(i.list)-1 {
		if !i.wasIteratorCalled {
			i.wasIteratorCalled = true
			return true
		}

		i.index++
		return true
	}

	return false
}
