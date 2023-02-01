package valuetypearray

import (
	"errors"
	t "linq/types"
)

type LinqArray[T interface {
	t.GenericType
}] struct {
	items []T
	pos   int
	size  int
}

func (la *LinqArray[T]) Init(size int) {
	la.size = size
	la.pos = 0
	la.items = make([]T, size)
}

func (la *LinqArray[T]) AddItem(t T) {
	la.items[la.pos] = t
	la.pos = la.pos + 1
}

func (la *LinqArray[T]) AddRange(list ...T) {
	for _, element := range list {
		la.items[la.pos] = element
		la.pos = la.pos + 1
	}
}

func (la *LinqArray[T]) RemoveITem(t T) error {
	index := la.FindIndexByValue(t)
	if index == -1 {
		return errors.New("ITEM_NOT_FOUND")
	}
	for i := index; i < la.pos-1; i++ {
		la.items[i] = la.items[i+1]
	}
	la.items[la.pos-1] = la.defaultValue()
	la.pos = la.pos - 1
	return nil
}

func (la *LinqArray[T]) RemoveDuplicates() {
	result := la.Distinct()
	la.Init(result.size)
	for _, value := range result.ToArray() {
		la.AddItem(value)
	}
}

func (la *LinqArray[T]) ToArray() []T {
	return la.items[:la.pos]
}

func (la *LinqArray[T]) Size() int {
	return la.pos
}

func (la *LinqArray[T]) Contains(item T) bool {
	for _, n := range la.ToArray() {
		if item == n {
			return true
		}
	}
	return false
}

func (la *LinqArray[T]) FindIndexByValue(item T) int {
	for index, n := range la.ToArray() {
		if item == n {
			return index
		}
	}
	return -1
}

func (la *LinqArray[T]) FindValueByIndex(index int) (T, error) {
	if index < 0 || index >= la.pos {
		return la.defaultValue(), errors.New("INDEX_NOT_FOUND")
	}
	return la.items[index], nil
}

func (la *LinqArray[T]) Where(f func(T) bool) *LinqArray[T] {
	filterMap := make(map[int]T)
	pos := 0
	for _, value := range la.ToArray() {
		if f(value) {
			filterMap[pos] = value
			pos = pos + 1
		}
	}
	var array LinqArray[T]
	array.Init(pos)
	for i := 0; i < pos; i++ {
		array.AddItem(filterMap[i])
	}
	return &array
}

func (la *LinqArray[T]) Distinct() *LinqArray[T] {
	filterMap := make(map[T]bool)
	pos := 0
	itemMap := make(map[int]T)
	for _, value := range la.ToArray() {
		if !filterMap[value] {
			filterMap[value] = true
			itemMap[pos] = value
			pos = pos + 1
		}
	}
	var array LinqArray[T]
	array.Init(pos)
	for i := 0; i < pos; i++ {
		array.AddItem(itemMap[i])
	}
	return &array
}

func (la *LinqArray[T]) GroupBy() map[T]*LinqArray[T] {
	filterMap := make(map[T][]T)
	pos := 0
	for _, value := range la.ToArray() {
		filterMap[value] = append(filterMap[value], value)
	}

	resultMap := make(map[T]*LinqArray[T])
	for i, value := range filterMap {
		var array LinqArray[T]
		pos = len(value)
		array.Init(pos)
		for i := 0; i < pos; i++ {
			array.AddItem(value[i])
		}
		resultMap[i] = &array
	}

	return resultMap
}

func (la *LinqArray[T]) defaultValue() T {
	var result T
	return result
}
