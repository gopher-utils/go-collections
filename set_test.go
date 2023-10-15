package collections_test

import (
	"testing"

	"github.com/gopher-utils/go-collections"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	intSetPointer := collections.NewSet[int]()
	stringSetPointer := collections.NewSet[string]()

	assert.NotNil(t, intSetPointer)
	assert.NotNil(t, stringSetPointer)
}

func TestToSet(t *testing.T) {
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
		{
			array: []int{1, 1, 2, 3, 3},
		},
	}

	for _, input := range inputs {
		s := collections.ToSet(input.array)
		for _, e := range input.array {
			assert.Equal(t, true, s.Contains(e))
		}
		// Checking less or equal because the ToSet removes duplicates
		assert.LessOrEqual(t, s.Size(), len(input.array))
	}
}

func TestSetAdd(t *testing.T) {
	s := collections.NewSet[int]()
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, e := range inputs {
		s.Add(e)
		assert.Equal(t, true, s.Contains(e))
	}
}

func TestSetContains(t *testing.T) {
	s := collections.ToSet([]int{1, 2, 2, 4, 6, 7, 9, 10})
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
		actualOutput := s.Contains(tc.input)
		assert.Equal(t, tc.output, actualOutput)
	}
}

func TestSetExtend(t *testing.T) {
	testCases := []struct {
		set1   *collections.Set[int]
		set2   *collections.Set[int]
		output *collections.Set[int]
	}{
		{
			set1:   collections.ToSet([]int{1, 2, 3, 5}),
			set2:   collections.ToSet([]int{3, 1, 4, 2, 1, 3, 3}),
			output: collections.ToSet([]int{1, 2, 3, 5, 4}),
		},
		{
			set1:   collections.ToSet([]int{1, 2, 3, 4, 5}),
			set2:   collections.ToSet([]int{6, 7, 8, 9, 10}),
			output: collections.ToSet([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
		{
			set1:   collections.ToSet([]int{1, 1, 1, 1, 1}),
			set2:   collections.ToSet([]int{0, 0}),
			output: collections.ToSet([]int{1, 1, 1, 1, 1, 0, 0}),
		},
	}

	for _, tc := range testCases {
		tc.set1.Extend(tc.set2)
		assert.Equal(t, tc.output, tc.set1)
	}
}

func TestSetClear(t *testing.T) {
	testCases := []struct {
		inputSet *collections.Set[int]
		output   *collections.Set[int]
	}{
		{
			inputSet: collections.ToSet([]int{1, 2, 3}),
			output:   collections.ToSet([]int{}),
		},
		{
			inputSet: collections.ToSet([]int{1}),
			output:   collections.ToSet([]int{}),
		},
		{
			inputSet: collections.ToSet([]int{}),
			output:   collections.ToSet([]int{}),
		},
	}

	for _, tc := range testCases {
		tc.inputSet.Clear()
		assert.Equal(t, tc.inputSet, tc.output)
	}
}

func TestSetRemove(t *testing.T) {
	testCases := []struct {
		inputSet     *collections.Set[int]
		inputElement int
		output       *collections.Set[int]
		err          error
	}{
		{
			inputSet:     collections.ToSet([]int{1, 2, 3}),
			inputElement: 1,
			output:       collections.ToSet([]int{2, 3}),
			err:          nil,
		},
		{
			inputSet:     collections.ToSet([]int{1}),
			inputElement: 1,
			output:       collections.ToSet([]int{}),
			err:          nil,
		},
		{
			inputSet:     collections.ToSet([]int{}),
			inputElement: 1,
			output:       collections.ToSet([]int{}),
			err:          collections.ErrItemNotFound,
		},
		{
			inputSet:     collections.ToSet([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 4,
			output:       collections.ToSet([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 2, 1, 3, 3}),
			err:          nil,
		},
		{
			inputSet:     collections.ToSet([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			inputElement: 10,
			output:       collections.ToSet([]int{1, 2, 2, 1, 3, 5, 5, 3, 3, 1, 4, 2, 1, 3, 3}),
			err:          collections.ErrItemNotFound,
		},
	}

	for _, tc := range testCases {
		err := tc.inputSet.Remove(tc.inputElement)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.inputSet, tc.output)
	}
}

func TestSetSize(t *testing.T) {
	testCases := []struct {
		inputSet *collections.Set[int]
		output   int
	}{
		{
			inputSet: collections.ToSet([]int{}),
			output:   0,
		},
		{
			inputSet: collections.ToSet([]int{1}),
			output:   1,
		},
		{
			inputSet: collections.ToSet([]int{1, 1, 2, 1, 2}),
			output:   2,
		},
		{
			inputSet: collections.ToSet([]int{1, 2, 3, 4, 5}),
			output:   5,
		},
	}

	for _, tc := range testCases {
		output := tc.inputSet.Size()
		assert.Equal(t, tc.output, output)
	}
}

func TestSetString(t *testing.T) {
	testCases := []struct {
		inputSet *collections.Set[int]
		output   string
	}{
		{
			inputSet: collections.ToSet([]int{}),
			output:   "{}",
		},
		{
			inputSet: collections.ToSet([]int{1}),
			output:   "{1}",
		},
	}

	for _, tc := range testCases {
		output := tc.inputSet.String()
		assert.Equal(t, tc.output, output)
	}
}

func TestSetToArray(t *testing.T) {
	testCases := []struct {
		inputSet *collections.Set[int]
		output   []int
	}{
		{
			inputSet: collections.ToSet([]int{}),
			output:   []int{},
		},
		{
			inputSet: collections.ToSet([]int{1}),
			output:   []int{1},
		},
		{
			inputSet: collections.ToSet([]int{1, 1, 1, 2}),
			output:   []int{1, 2},
		},
	}

	for _, tc := range testCases {
		output := tc.inputSet.ToArray()
		assert.Equal(t, len(tc.output), len(output))
	}
}

func TestSetType(t *testing.T) {
	s := collections.NewSet[int]()

	assert.Equal(t, collections.TypeSet, s.Type())
}

func TestSetUnion(t *testing.T) {
	testCases := []struct {
		set1   *collections.Set[int]
		set2   *collections.Set[int]
		output *collections.Set[int]
	}{
		{
			set1:   collections.ToSet([]int{1, 2, 3}),
			set2:   collections.ToSet([]int{4, 5, 6}),
			output: collections.ToSet([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			set1:   collections.ToSet([]int{1, 2, 3}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{1, 2, 3, 4, 5}),
		},
		{
			set1:   collections.ToSet([]int{}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{3, 4, 5}),
		},
		{
			set1:   collections.ToSet([]int{1, 2}),
			set2:   collections.ToSet([]int{}),
			output: collections.ToSet([]int{1, 2}),
		},
	}

	for _, tc := range testCases {
		s := tc.set1.Union(tc.set2)
		assert.Equal(t, tc.output, s)
	}
}

func TestSetIntersection(t *testing.T) {
	testCases := []struct {
		set1   *collections.Set[int]
		set2   *collections.Set[int]
		output *collections.Set[int]
	}{
		{
			set1:   collections.ToSet([]int{1, 2, 3}),
			set2:   collections.ToSet([]int{4, 5, 6}),
			output: collections.ToSet([]int{}),
		},
		{
			set1:   collections.ToSet([]int{4, 1, 2, 3}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{3, 4}),
		},
		{
			set1:   collections.ToSet([]int{}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{}),
		},
		{
			set1:   collections.ToSet([]int{1}),
			set2:   collections.ToSet([]int{1, 2}),
			output: collections.ToSet([]int{1}),
		},
	}

	for _, tc := range testCases {
		s := tc.set1.Intersection(tc.set2)
		assert.Equal(t, tc.output, s)
	}
}

func TestSetDifference(t *testing.T) {
	testCases := []struct {
		set1   *collections.Set[int]
		set2   *collections.Set[int]
		output *collections.Set[int]
	}{
		{
			set1:   collections.ToSet([]int{1, 2, 3}),
			set2:   collections.ToSet([]int{4, 5, 6}),
			output: collections.ToSet([]int{1, 2, 3}),
		},
		{
			set1:   collections.ToSet([]int{1, 2, 3}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{1, 2}),
		},
		{
			set1:   collections.ToSet([]int{}),
			set2:   collections.ToSet([]int{3, 4, 5}),
			output: collections.ToSet([]int{}),
		},
		{
			set1:   collections.ToSet([]int{1, 2}),
			set2:   collections.ToSet([]int{}),
			output: collections.ToSet([]int{1, 2}),
		},
	}

	for _, tc := range testCases {
		s := tc.set1.Difference(tc.set2)
		assert.Equal(t, tc.output, s)
	}
}
