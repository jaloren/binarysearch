package tree

type node[T any] struct {
	value  Sorter[T]
	left   *node[T]
	right  *node[T] // larger
	parent *node[T]
}

type Tree[T any] struct {
	root *node[T]
	size int
}

func New[T any](root Sorter[T]) *Tree[T] {
	return &Tree[T]{
		root: &node[T]{
			value: root,
		},
		size: 1,
	}
}

func (t *Tree[T]) Insert(value T) {
	current := t.root
	for {

	}

}
