package collections_test

import (
	"testing"

	"github.com/gopher-utils/go-collections"

	"github.com/stretchr/testify/assert"
)

func TestAvg(t *testing.T) {
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
			output:    3,
		},
	}

	for _, tc := range testCases {
		output := collections.Avg[int](tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestMax(t *testing.T) {
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
		output := collections.Max[int](tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestMin(t *testing.T) {
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
			output:    1,
		},
		{
			inputList: collections.ToList([]int{5, 4, 3, 2, 1, 0}),
			output:    0,
		},
	}

	for _, tc := range testCases {
		output := collections.Min[int](tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestSum(t *testing.T) {
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
			output:    15,
		},
	}

	for _, tc := range testCases {
		output := collections.Sum[int](tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}
