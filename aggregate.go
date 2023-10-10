package collections

// Returns the average of all elements present in the given collection.
// This method supports only numerical types.
func Avg[V Number](c Collection[V]) V {
	switch c.Type() {
	case TypeList:
		size := V(c.Size())
		if size == 0 {
			return 0
		}
		return Sum[V](c) / size
	}
	panic("unknown collection type")
}

// Returns maximum element present in the given collection.
// This method supports only numerical types.
func Max[V Number](c Collection[V]) V {
	switch c.Type() {
	case TypeList:
		l := c.(List[V])
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
func Min[V Number](c Collection[V]) V {
	switch c.Type() {
	case TypeList:
		l := c.(List[V])
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
func Sum[V Number](c Collection[V]) V {
	switch c.Type() {
	case TypeList:
		l := c.(List[V])
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
