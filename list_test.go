package collections_test

import (
	"testing"

	"github.com/gopher-utils/go-collections"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	intListPointer := collections.NewList[int](1)
	stringListPointer := collections.NewList[string](1)

	assert.NotNil(t, intListPointer)
	assert.NotNil(t, stringListPointer)
}

func BenchmarkNewList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = collections.NewList[int](10)
	}
}

func TestToList(t *testing.T) {
	inputs := []struct {
		array []int
	}{
		{
			array: []int{1, 2, 3},
		},
		{
			array: []int{1, 2, 3, 4, 5},
		},
		{
			array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, input := range inputs {
		l := collections.ToList(input.array)
		assert.Equal(t, input.array, l.ToArray())
	}
}

func BenchmarkToList(b *testing.B) {
	example := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_ = collections.ToList(example)
	}
}

func TestRepeatingList(t *testing.T) {
	inputs := []struct {
		element int
		times   int
		result  []int
	}{
		{
			element: 0,
			times:   0,
			result:  []int{},
		},
		{
			element: 0,
			times:   4,
			result:  []int{0, 0, 0, 0},
		},
		{
			element: 4,
			times:   0,
			result:  []int{},
		},
		{
			element: 1,
			times:   2,
			result:  []int{1, 1},
		},
	}

	for _, input := range inputs {
		l := collections.RepeatingList(input.element, input.times)
		assert.Equal(t, input.result, l.ToArray())
	}
}

func BenchmarkRepeatingList1_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = collections.RepeatingList(1, 10)
	}
}

func TestListAdd(t *testing.T) {
	l := collections.NewList[int](10)
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, input := range inputs {
		l.Add(input)
		assert.Equal(t, inputs[0:index+1], l.ToArray())
	}
}

func BenchmarkListAdd(b *testing.B) {
	l := collections.NewList[int](10)
	for i := 0; i < b.N; i++ {
		l.Add(1)
	}
}

func TestListContains(t *testing.T) {
	l := collections.ToList([]int{1, 2, 2, 4, 6, 7, 9, 10})
	testCases := []struct {
		input  int
		output bool
	}{
		{input: 1, output: true},
		{input: 2, output: true},
		{input: 3, output: false},
		{input: 4, output: true},
		{input: 5, output: false},
		{input: 6, output: true},
		{input: 7, output: true},
		{input: 8, output: false},
		{input: 9, output: true},
		{input: 10, output: true},
	}

	for _, tc := range testCases {
		actualOutput := l.Contains(tc.input)
		assert.Equal(t, tc.output, actualOutput)
	}
}

func BenchmarkListContainsTrue(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 4, 6, 7, 9, 10})
	for i := 0; i < b.N; i++ {
		l.Contains(6)
	}
}

func BenchmarkListContainsFalse(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 4, 6, 7, 9, 10})
	for i := 0; i < b.N; i++ {
		l.Contains(5)
	}
}

func TestListCountOf(t *testing.T) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	testCases := []struct {
		input  int
		output int
	}{
		{input: 1, output: 4},
		{input: 2, output: 3},
		{input: 3, output: 5},
		{input: 4, output: 1},
		{input: 5, output: 2},
	}

	for _, tc := range testCases {
		actualOutput := l.CountOf(tc.input)
		assert.Equal(t, tc.output, actualOutput)
	}
}

func BenchmarkListCountOf(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	for i := 0; i < b.N; i++ {
		l.CountOf(3)
	}
}

func TestListDistinct(t *testing.T) {
	testCases := []struct {
		input  *collections.List[int]
		output *collections.List[int]
	}{
		{
			input:  collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.ToList([]int{1, 2, 3, 5, 4}),
		},
		{
			input:  collections.ToList([]int{1, 1, 2, 2, 3, 3, 4, 5, 4}),
			output: collections.ToList([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.ToList([]int{1, 2, 3, 4, 5}),
			output: collections.ToList([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.ToList([]int{5, 4, 3, 2, 1}),
			output: collections.ToList([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		actualOutput := tc.input.Distinct()
		assert.Equal(t, tc.output, actualOutput)
	}
}

func BenchmarkListDistinct(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	for i := 0; i < b.N; i++ {
		l.Distinct()
	}
}

func TestListExtend(t *testing.T) {
	testCases := []struct {
		list1  *collections.List[int]
		list2  *collections.List[int]
		output *collections.List[int]
	}{
		{
			list1:  collections.ToList([]int{1, 2, 2, 1, 3, 5, 5}),
			list2:  collections.ToList([]int{3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
		},
		{
			list1:  collections.ToList([]int{1, 2, 3, 4, 5}),
			list2:  collections.ToList([]int{6, 7, 8, 9, 10}),
			output: collections.ToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
		{
			list1:  collections.ToList([]int{1, 1, 1, 1, 1}),
			list2:  collections.ToList([]int{0, 0}),
			output: collections.ToList([]int{1, 1, 1, 1, 1, 0, 0}),
		},
	}

	for _, tc := range testCases {
		tc.list1.Extend(tc.list2)
		assert.Equal(t, tc.output, tc.list1)
	}
}

func BenchmarkListExtend(b *testing.B) {
	list1 := collections.ToList([]int{1, 1, 1, 1, 1})
	list2 := collections.ToList([]int{0, 0})
	for i := 0; i < b.N; i++ {
		list1.Extend(list2)
	}
}

func TestListGet(t *testing.T) {
	l := collections.ToList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	testCases := []struct {
		input  int
		output int
		err    error
	}{
		{input: -1, output: 0, err: collections.ErrIndexOutOfRange},
		{input: 0, output: 0, err: nil},
		{input: 1, output: 1, err: nil},
		{input: 2, output: 2, err: nil},
		{input: 3, output: 3, err: nil},
		{input: 4, output: 4, err: nil},
		{input: 5, output: 5, err: nil},
		{input: 6, output: 6, err: nil},
		{input: 7, output: 7, err: nil},
		{input: 8, output: 8, err: nil},
		{input: 9, output: 9, err: nil},
		{input: 10, output: 10, err: nil},
		{input: 11, output: 0, err: collections.ErrIndexOutOfRange},
	}

	for _, tc := range testCases {
		output, err := l.Get(tc.input)
		assert.Equal(t, tc.output, output)
		assert.Equal(t, tc.err, err)
	}
}

func BenchmarkListGetFound(b *testing.B) {
	l := collections.ToList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		l.Get(5)
	}
}

func BenchmarkListGetNotFound(b *testing.B) {
	l := collections.ToList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		l.Get(-1)
	}
}

func TestIndexOf(t *testing.T) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	testCases := []struct {
		input  int
		output int
	}{
		{input: 0, output: -1},
		{input: 1, output: 0},
		{input: 2, output: 1},
		{input: 3, output: 2},
		{input: 4, output: 3},
		{input: 5, output: 4},
		{input: 6, output: -1},
	}

	for _, tc := range testCases {
		output := l.IndexOf(tc.input)
		assert.Equal(t, tc.output, output)
	}
}

func BenchmarkListIndexOf(b *testing.B) {
	l := collections.ToList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < b.N; i++ {
		l.IndexOf(5)
	}
}

func TestListRemoveAll(t *testing.T) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	testCases := []struct {
		input int
		err   error
	}{
		{input: 1, err: nil},
		{input: 2, err: nil},
		{input: 3, err: nil},
		{input: 4, err: nil},
		{input: 5, err: nil},
		{input: 1, err: collections.ErrItemNotFound},
	}

	for _, tc := range testCases {
		err := l.RemoveAll(tc.input)
		removed := !l.Contains(tc.input)
		assert.Equal(t, true, removed)
		assert.Equal(t, tc.err, err)
	}
}

func BenchmarkListRemoveAll(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	for i := 0; i < b.N; i++ {
		l.RemoveAll(5)
	}
}

func TestListRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input  *collections.List[int]
		output *collections.List[int]
	}{
		{
			input:  collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			output: collections.ToList([]int{1, 2, 3, 5, 4}),
		},
		{
			input:  collections.ToList([]int{1, 1, 2, 2, 3, 3, 4, 5, 4}),
			output: collections.ToList([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.ToList([]int{1, 2, 3, 4, 5}),
			output: collections.ToList([]int{1, 2, 3, 4, 5}),
		},
		{
			input:  collections.ToList([]int{5, 4, 3, 2, 1}),
			output: collections.ToList([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tc := range testCases {
		tc.input.RemoveDuplicates()
		assert.Equal(t, tc.input.ToArray(), tc.output.ToArray())
	}
}

func BenchmarkListRemoveDuplicates(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	for i := 0; i < b.N; i++ {
		l.RemoveDuplicates()
	}
}

func TestListRemoveFirst(t *testing.T) {
	testCases := []struct {
		inputList    *collections.List[int]
		inputElement int
		output       *collections.List[int]
		err          error
	}{
		{
			inputList:    collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 1,
			output:       collections.ToList([]int{2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			err:          nil,
		},
		{
			inputList:    collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 4,
			output:       collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 2, 1, 3, 3}),
			err:          nil,
		},
		{
			inputList:    collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 10,
			output:       collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			err:          collections.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		err := tc.inputList.RemoveFirst(tc.inputElement)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.inputList.ToArray(), tc.output.ToArray())
	}
}

func BenchmarkListRemoveFirst(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	for i := 0; i < b.N; i++ {
		l.RemoveFirst(1)
	}
}

func TestListWhere(t *testing.T) {
	testCases := []struct {
		inputList  *collections.List[int]
		outputList *collections.List[int]
		filterFunc func(int) bool
	}{
		{
			inputList:  collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			outputList: collections.ToList([]int{2, 2, 4, 2}),
			filterFunc: func(i int) bool { return i%2 == 0 },
		},
		{
			inputList:  collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			outputList: collections.ToList([]int{1, 2, 2, 1, 1, 2, 1}),
			filterFunc: func(i int) bool { return i < 3 },
		},
	}

	for _, tc := range testCases {
		outputList := tc.inputList.Where(tc.filterFunc)
		assert.Equal(t, tc.outputList.ToArray(), outputList.ToArray())
	}
}

func BenchmarkListWhere(b *testing.B) {
	l := collections.ToList([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3})
	filterFunc := func(i int) bool { return i%2 == 0 }
	for i := 0; i < b.N; i++ {
		l.Where(filterFunc)
	}
}

func TestListSize(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[int]
		output    int
	}{
		{
			inputList: collections.ToList([]int{}),
			output:    0,
		},
		{
			inputList: collections.ToList([]int{1}),
			output:    1,
		},
		{
			inputList: collections.ToList([]int{1, 2, 3, 4, 5}),
			output:    5,
		},
	}

	for _, tc := range testCases {
		output := tc.inputList.Size()
		assert.Equal(t, tc.output, output)
	}
}

func BenchmarkListSize(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	for i := 0; i < b.N; i++ {
		l.Size()
	}
}

func TestListString(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[int]
		output    string
	}{
		{
			inputList: collections.ToList([]int{}),
			output:    "[]",
		},
		{
			inputList: collections.ToList([]int{1}),
			output:    "[1]",
		},
		{
			inputList: collections.ToList([]int{1, 2, 3, 4, 5}),
			output:    "[1,2,3,4,5]",
		},
	}

	for _, tc := range testCases {
		output := tc.inputList.String()
		assert.Equal(t, tc.output, output)
	}
}

func BenchmarkListString(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	for i := 0; i < b.N; i++ {
		l.String()
	}
}

func TestListType(t *testing.T) {
	l := collections.NewList[int](0)

	assert.Equal(t, collections.TypeList, l.Type())
}

func BenchmarkListType(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	for i := 0; i < b.N; i++ {
		l.Type()
	}
}

func TestListMap(t *testing.T) {
	studentMarks := []struct {
		name  string
		marks int
	}{
		{name: "Jack", marks: 100},
		{name: "Jill", marks: 90},
		{name: "Tom", marks: 75},
		{name: "Peter", marks: 50},
	}

	expected := collections.ToList([]int{100, 90, 75, 50})

	l := collections.ToList(studentMarks)

	getMarksCallbackFunc := func(elem struct {
		name  string
		marks int
	}, index int) int {
		return elem.marks
	}

	nl := collections.Map(l, getMarksCallbackFunc)

	assert.Equal(t, expected, nl)
}

func BenchmarkListMap(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	funcMap := func(e int, i int) int { return e * e }
	for i := 0; i < b.N; i++ {
		collections.Map(l, funcMap)
	}
}

func TestListReduce(t *testing.T) {
	testCases := []struct {
		inputList    *collections.List[int]
		callbackFunc func(result int, current int) int
		initialValue int
		output       int
	}{
		{
			inputList:    collections.ToList([]int{1, 2, 3, 4, 5}),
			callbackFunc: func(result int, current int) int { return result + current },
			initialValue: 0,
			output:       15,
		},
		{
			inputList:    collections.ToList([]int{1, 2, 3, 4, 5}),
			callbackFunc: func(result int, current int) int { return result * current },
			initialValue: 1,
			output:       120,
		},
		{
			inputList:    collections.ToList([]int{1, 2, 3, 4, 5}),
			callbackFunc: func(result int, current int) int { return result + (current * current) },
			initialValue: 0,
			output:       55,
		},
	}

	for _, tc := range testCases {
		output := collections.Reduce(tc.inputList, tc.callbackFunc, tc.initialValue)
		assert.Equal(t, tc.output, output)
	}
}

func BenchmarkListReduce(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	funcMap := func(e int, i int) int { return e * e }
	for i := 0; i < b.N; i++ {
		collections.Reduce(l, funcMap, 0)
	}
}

func TestSortIntegers(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[int]
		output    *collections.List[int]
	}{
		{
			inputList: collections.ToList([]int{5, 4, 3, 2, 1}),
			output:    collections.ToList([]int{1, 2, 3, 4, 5}),
		},
		{
			inputList: collections.ToList([]int{21, 13, 45, 44, 21}),
			output:    collections.ToList([]int{13, 21, 21, 44, 45}),
		},
		{
			inputList: collections.ToList([]int{1}),
			output:    collections.ToList([]int{1}),
		},
	}

	for _, tc := range testCases {
		tc.inputList.Sort(func(a, b int) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		})
		assert.Equal(t, tc.output, tc.inputList)
	}
}

func BenchmarkListSortIntegers(b *testing.B) {
	l := collections.ToList([]int{1, 2, 3, 4, 5})
	funcSort := func(a, b int) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}
	for i := 0; i < b.N; i++ {
		l.Sort(funcSort)
	}
}

func TestSortStrings(t *testing.T) {
	testCases := []struct {
		inputList *collections.List[string]
		output    *collections.List[string]
	}{
		{
			inputList: collections.ToList([]string{"e", "d", "c", "b", "a"}),
			output:    collections.ToList([]string{"a", "b", "c", "d", "e"}),
		},
		{
			inputList: collections.ToList([]string{"banana", "apple", "cherry", "date"}),
			output:    collections.ToList([]string{"apple", "banana", "cherry", "date"}),
		},
	}

	for _, tc := range testCases {
		tc.inputList.Sort(func(a, b string) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		})
		assert.Equal(t, tc.output, tc.inputList)
	}

}

func BenchmarkListSortStrings(b *testing.B) {
	l := collections.ToList([]string{"e", "d", "c", "b", "a"})
	funcSort := func(a, b string) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}
	for i := 0; i < b.N; i++ {
		l.Sort(funcSort)
	}
}
