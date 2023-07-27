package collections

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrIndexOutOfRange = errors.New("linq: index out of range")
	ErrItemNotFound    = errors.New("linq: item not found")
)

// Collection that stores homogenous elements in a fixed order.
type List[T CollectionElement] struct {
	items []T
}

// Factory method to create an empty list with predefined capacity.
func NewEmptyList[T CollectionElement](capacity int) *List[T] {
	l := &List[T]{}
	l.items = make([]T, 0, capacity)
	return l
}

// Factory method to create a list from an array.
func NewListFromArray[T CollectionElement](array []T) *List[T] {
	l := &List[T]{}
	l.items = make([]T, 0, len(array))
	l.items = append(l.items, array...)
	return l
}

// Factory method to create a list with repeating values
func NewRepeatedList[T CollectionElement](element T, times int) *List[T] {
	l := &List[T]{}
	l.items = make([]T, 0, times)
	for i := 0; i < times; i++ {
		l.Add(element)
	}
	return l
}

// Add element to list.
func (l *List[T]) Add(item T) {
	l.items = append(l.items, item)
}

// Checks whether an element is present in the list.
func (l *List[T]) Contains(item T) bool {
	return l.IndexOf(item) != -1
}

// Returns number of occurences of given element in the list.
func (l *List[T]) CountOf(item T) (count int) {
	for _, e := range l.items {
		if e == item {
			count++
		}
	}
	return count
}

// Returns a new list containing unique elements.
func (l *List[T]) Distinct() *List[T] {
	filterMap := make(map[T]bool)
	distinctList := NewEmptyList[T](len(l.items))

	for _, e := range l.items {
		if !filterMap[e] {
			distinctList.Add(e)
			filterMap[e] = true
		}
	}

	return distinctList
}

// Concatenate a list with another list.
func (l *List[T]) Extend(l2 *List[T]) {
	l.items = append(l.items, l2.items...)
}

// Returns list element for a valid index.
// Returns error for an invalid index.
func (l *List[T]) Get(index int) (item T, err error) {
	if index >= 0 && index < l.Size() {
		item = l.items[index]
		return item, nil
	}

	err = ErrIndexOutOfRange
	return item, err
}

// Returns the index of the first occurence of the element.
// If the element is not present in the list, it returns -1.
func (l *List[T]) IndexOf(item T) int {
	for i, e := range l.items {
		if e == item {
			return i
		}
	}
	return -1
}

// Removes all occurences of the given element from the list.
// Returns an error if the element is not present in the list.
func (l *List[T]) RemoveAll(item T) error {

	if count := l.CountOf(item); count == 0 {
		return ErrItemNotFound
	} else if count == 1 {
		return l.RemoveFirst(item)
	}

	var updatedItems []T
	for _, e := range l.items {
		if e != item {
			updatedItems = append(updatedItems, e)
		}
	}
	l.items = updatedItems
	return nil
}

// Removes duplicates of elements in the list.
func (l *List[T]) RemoveDuplicates() {
	l.items = l.Distinct().items
}

// Removes first occurence of the given element from the list.
// Returns an error if the element is not present in the list.
func (l *List[T]) RemoveFirst(item T) error {
	var index int
	if index = l.IndexOf(item); index == -1 {
		return ErrItemNotFound
	}
	l.items = append(l.items[:index], l.items[index+1:]...)
	return nil
}

// Returns an slice containing elements of the list.
func (l *List[T]) ToArray() []T {
	return l.items
}

// Returns a filtered list based on the provided boolean function f.
func (l *List[T]) Where(f func(T) bool) *List[T] {
	resultList := NewEmptyList[T](len(l.items))

	for _, e := range l.items {
		if f(e) {
			resultList.Add(e)
		}
	}

	return resultList
}

// Returns the number of elements in the list.
func (l List[T]) Size() int {
	return len(l.items)
}

// Returns a string description of the list.
func (l List[T]) String() string {
	resultStrings := make([]string, 0, l.Size())
	for _, e := range l.items {
		resultStrings = append(resultStrings, fmt.Sprint(e))
	}

	return "[" + strings.Join(resultStrings, ",") + "]"
}

func (_ List[T]) Type() CollectionType {
	return TypeList
}

// Use Map method to transform a list of a given type to another type.
func Map[T CollectionElement, E CollectionElement](l *List[T], callback listMapFunction[T, E]) *List[E] {
	result := NewEmptyList[E](l.Size())
	for i, e := range l.items {
		result.Add(callback(e, i))
	}
	return result
}

// Use Reduce to reduce the given list elements to a single element of same type T based on a callback function.
func Reduce[T CollectionElement](l *List[T], callback listReduceFunction[T], initialValue T) T {
	result := initialValue
	for _, e := range l.items {
		result = callback(result, e)
	}
	return result
}

// Type of callback function that needs to be passed to Map method.
type listMapFunction[T CollectionElement, E CollectionElement] func(element T, index int) E

// Type of callback function that needs to be passed to Reduce method.
type listReduceFunction[T CollectionElement] func(result T, item T) T
