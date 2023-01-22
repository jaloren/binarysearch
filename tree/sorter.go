package tree

import "golang.org/x/exp/constraints"

type Comparer[T any] interface {
	compareTo(x T, y T) int
}

type OrderedType[T constraints.Ordered] struct {
}

func (o OrderedType[T]) compareTo(x T, y T) int {
	if x < y {
		return -1
	} else if x > y {
		return 1
	} else {
		return 0
	}
}

func NewComparer[T constraints.Ordered]() OrderedType[T] {
	return OrderedType[T]{}
}
