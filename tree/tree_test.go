package tree

import (
	"fmt"
	"testing"
)

func TestSorter(t *testing.T) {
	var s Sorter[int]
	s = NewComparer(45)
	fmt.Println(howBig(s, 10))
}

func howBig[T any](s Sorter[T], input T) int {
	return s.compareTo(input)
}
