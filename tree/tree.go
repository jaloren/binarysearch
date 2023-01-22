package tree

import "fmt"

type Node[T any] struct {
	value  T
	left   *Node[T] // smaller
	right  *Node[T] // larger
	parent *Node[T]
}

type Tree[T any] struct {
	root     *Node[T]
	comparer Comparer[T]
	size     int
}

func New[T any](value T, comparer Comparer[T]) *Tree[T] {
	return &Tree[T]{
		root: &Node[T]{
			value: value,
		},
		comparer: comparer,
		size:     1,
	}
}

func (t *Tree[T]) Insert(value T) {
	current := t.root
	for {
		result := t.comparer.compareTo(t.root.value, value)
		switch result {
		case 0:
			return
		case 1:
			if current.right == nil {
				current.right = &Node[T]{
					value:  value,
					parent: current,
				}
				return
			} else {
				current = current.right
				continue
			}
		case -1:
			if current.left == nil {
				current.left = &Node[T]{
					value:  value,
					parent: current,
				}
				return
			} else {
				current = current.left
				continue
			}
		default:
			panic(fmt.Sprintf("comparison function return %d expected 0, 1 or -1", result))
		}
	}
}

type Iterator[T any] struct {
	position *Node[T]
}

func (i *Iterator[T]) Iter() (T, bool) {
	current := i.position
	if current == nil {
		var empty T
		return empty, false
	}
	if current.left != nil {
		i.position = current.left
		return current.value, true
	} else if current.right != nil {
		i.position = current.right
		return current.value, true
	} else {
		i.position = nil
		return current.value, true
	}
}

func (t *Tree[T]) NewIterator() *Iterator[T] {
	return &Iterator[T]{position: t.root}
}
