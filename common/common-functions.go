package common

import (
	sort "linq/sort"
	ty "linq/types"
	"linq/valuetypearray"
)

func OrderBy[V ty.Number](la *valuetypearray.LinqArray[V]) *valuetypearray.LinqArray[V] {
	result := sort.Init(la.ToArray())
	var array valuetypearray.LinqArray[V]
	array.Init(la.Size())
	for _, value := range result {
		array.AddItem(value)
	}
	return &array
}

func OrderByDesc[V ty.Number](la *valuetypearray.LinqArray[V]) *valuetypearray.LinqArray[V] {
	result := sort.Init(la.ToArray())
	var array valuetypearray.LinqArray[V]
	array.Init(la.Size())
	for i := la.Size() - 1; i > 0; i-- {
		array.AddItem(result[i])
	}
	return &array
}

func Max[V ty.Number](la *valuetypearray.LinqArray[V]) V {
	items := la.ToArray()
	max := items[0]
	for i := 1; i < la.Size(); i++ {
		if max < items[i] {
			max = items[i]
		}
	}
	return max
}

func Min[V ty.Number](la *valuetypearray.LinqArray[V]) V {
	items := la.ToArray()
	min := items[0]
	for i := 1; i < la.Size(); i++ {
		if items[i] < min {
			min = items[i]
		}
	}
	return min
}

func Sum[V ty.Number](la *valuetypearray.LinqArray[V]) V {
	items := la.ToArray()
	sum := items[0]
	for i := 1; i < la.Size(); i++ {
		sum = sum + items[i]
	}
	return sum
}

func Avg[V ty.Number](la *valuetypearray.LinqArray[V]) V {
	sum := Sum(la)
	size := la.Size()

	return sum / V(size)
}
