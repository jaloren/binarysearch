package tree

import (
	"fmt"
	"testing"
)

func TestSorter(t *testing.T) {
	tree := New[int](1, NewComparer[int]())
	tree.Insert(5)
	tree.Insert(253)
	tree.Insert(233)
	tree.Insert(1000)
	tree.Insert(10)

	iter := tree.NewIterator()
	for {
		val, ok := iter.Iter()
		if !ok {
			break
		} else {
			fmt.Println(val)
		}
	}

}
