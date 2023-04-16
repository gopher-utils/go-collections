package collections

// Returns maximum element present in the given collection.
// This method supports only numerical types.
func Max[V Number](c Collection[V]) V {
	switch c.Type() {
	case LIST:
		l := c.(List[V])
		items := l.ToArray()
		max := items[0]
		for i := 1; i < l.Size(); i++ {
			if max < items[i] {
				max = items[i]
			}
		}
		return max
	case SET:
		panic("not implemented")
	}
	panic("unknown collection type")
}

// Returns minimum element present in the given collection.
// This method supports only numerical types.
func Min[V Number](c Collection[V]) V {
	switch c.Type() {
	case LIST:
		l := c.(List[V])
		items := l.ToArray()
		min := items[0]
		for i := 1; i < l.Size(); i++ {
			if min > items[i] {
				min = items[i]
			}
		}
		return min
	case SET:
		panic("not implemented")
	}
	panic("unknown collection type")
}

// Returns the sum of all elements present in the given collection.
// This method supports only numerical types.
func Sum[V Number](c Collection[V]) V {
	switch c.Type() {
	case LIST:
		l := c.(List[V])
		items := l.ToArray()
		sum := items[0]
		for i := 1; i < l.Size(); i++ {
			sum += items[i]
		}
		return sum
	case SET:
		panic("not implemented")
	}
	panic("unknown collection type")
}

// Returns the average of all elements present in the given collection.
// This method supports only numerical types.
func Avg[V Number](c Collection[V]) V {
	switch c.Type() {
	case LIST:
		return Sum(c) / V(c.Size())
	case SET:
		panic("not implemented")
	}
	panic("unknown collection type")
}
