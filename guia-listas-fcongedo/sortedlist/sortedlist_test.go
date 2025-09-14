package sortedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedList(t *testing.T) {
	sl := NewSortedList[int]()

	sl.Insert(3)
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(-4)
	sl.Insert(3)
	sl.Insert(0)
	sl.Insert(5)

	nodo := sl.lista.Head()
	for _, v := range []int{-4, 0, 1, 2, 3, 3, 5} {
		assert.Equal(t, v, nodo.Data())
		nodo = nodo.Next()
	}

}
