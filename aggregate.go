package collections

// Returns the average of all elements present in the given collection.
// This method supports only numerical types.
func Avg[V Number](c Collection[V]) V {
	size := V(c.Size())
	if size == 0 {
		return 0
	}
	return Sum[V](c) / size
}

// Returns maximum element present in the given collection.
// This method supports only numerical types.
func Max[V Number](c Collection[V]) V {
	slice := toSlice[V](c)
	size := len(slice)
	if size == 0 {
		return 0
	}
	max := slice[0]
	for i := 1; i < size; i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}

	return max
}

// Returns minimum element present in the given collection.
// This method supports only numerical types.
func Min[V Number](c Collection[V]) V {
	slice := toSlice[V](c)
	size := len(slice)
	if size == 0 {
		return 0
	}
	min := slice[0]
	for i := 1; i < size; i++ {
		if min > slice[i] {
			min = slice[i]
		}
	}

	return min
}

// Returns the sum of all elements present in the given collection.
// This method supports only numerical types.
func Sum[V Number](c Collection[V]) V {
	slice := toSlice[V](c)
	sum := V(0)
	for _, elem := range slice {
		sum += elem
	}
	return sum
}

func toSlice[V Number](c Collection[V]) []V {
	switch c.Type() {
	case TypeList:
		return c.(*List[V]).ToArray()
	case TypeSet:
		return c.(*Set[V]).ToArray()
	}
	panic("unknown collection type")
}
