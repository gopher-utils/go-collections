package math

import (
	"github.com/gopinathr143/golang-linq/collections"
	"github.com/gopinathr143/golang-linq/collections/list"
)

// Returns the average of all elements present in the given collection.
// This method supports only numerical types.
func Avg[V collections.Number](c collections.Collection[V]) V {
	switch c.Type() {
	case collections.TypeList:
		size := V(c.Size())
		if size == 0 {
			return 0
		}
		return Sum(c) / size
	}
	panic("unknown collection type")
}

// Returns maximum element present in the given collection.
// This method supports only numerical types.
func Max[V collections.Number](c collections.Collection[V]) V {
	switch c.Type() {
	case collections.TypeList:
		l := c.(list.List[V])
		if l.Size() == 0 {
			return 0
		}
		items := l.ToArray()
		max := items[0]
		for i := 1; i < l.Size(); i++ {
			if max < items[i] {
				max = items[i]
			}
		}
		return max
	}
	panic("unknown collection type")
}

// Returns minimum element present in the given collection.
// This method supports only numerical types.
func Min[V collections.Number](c collections.Collection[V]) V {
	switch c.Type() {
	case collections.TypeList:
		l := c.(list.List[V])
		if l.Size() == 0 {
			return 0
		}
		items := l.ToArray()
		min := items[0]
		for i := 1; i < l.Size(); i++ {
			if min > items[i] {
				min = items[i]
			}
		}
		return min
	}
	panic("unknown collection type")
}

// Returns the sum of all elements present in the given collection.
// This method supports only numerical types.
func Sum[V collections.Number](c collections.Collection[V]) V {
	switch c.Type() {
	case collections.TypeList:
		l := c.(list.List[V])
		if l.Size() == 0 {
			return 0
		}
		items := l.ToArray()
		sum := items[0]
		for i := 1; i < l.Size(); i++ {
			sum += items[i]
		}
		return sum
	}
	panic("unknown collection type")
}
