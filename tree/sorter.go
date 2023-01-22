package tree

import "golang.org/x/exp/constraints"

type Sorter[T any] interface {
	compareTo(x T) int
}

type OrderedType[T constraints.Ordered] struct {
	value T
}

func (o OrderedType[T]) compareTo(x T) int {
	if o.value < x {
		return -1
	} else if o.value > x {
		return 1
	} else {
		return 0
	}
}

func NewComparer[T constraints.Ordered](value T) OrderedType[T] {
	return OrderedType[T]{value: value}
}
