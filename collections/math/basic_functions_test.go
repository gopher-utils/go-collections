package math_test

import (
	"testing"

	"github.com/gopinathr143/golang-linq/collections/list"
	"github.com/gopinathr143/golang-linq/collections/math"

	"github.com/stretchr/testify/assert"
)

func TestAvg(t *testing.T) {
	testCases := []struct {
		inputList *list.List[int]
		output    int
	}{
		{
			inputList: list.From([]int{}),
			output:    0,
		},
		{
			inputList: list.From([]int{1}),
			output:    1,
		},
		{
			inputList: list.From([]int{1, 2, 3, 4, 5}),
			output:    3,
		},
	}

	for _, tc := range testCases {
		output := math.Avg[int](*tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		inputList *list.List[int]
		output    int
	}{
		{
			inputList: list.From([]int{}),
			output:    0,
		},
		{
			inputList: list.From([]int{1}),
			output:    1,
		},
		{
			inputList: list.From([]int{1, 2, 3, 4, 5}),
			output:    5,
		},
	}

	for _, tc := range testCases {
		output := math.Max[int](*tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		inputList *list.List[int]
		output    int
	}{
		{
			inputList: list.From([]int{}),
			output:    0,
		},
		{
			inputList: list.From([]int{1}),
			output:    1,
		},
		{
			inputList: list.From([]int{1, 2, 3, 4, 5}),
			output:    1,
		},
		{
			inputList: list.From([]int{5, 4, 3, 2, 1, 0}),
			output:    0,
		},
	}

	for _, tc := range testCases {
		output := math.Min[int](*tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}

func TestSum(t *testing.T) {
	testCases := []struct {
		inputList *list.List[int]
		output    int
	}{
		{
			inputList: list.From([]int{}),
			output:    0,
		},
		{
			inputList: list.From([]int{1}),
			output:    1,
		},
		{
			inputList: list.From([]int{1, 2, 3, 4, 5}),
			output:    15,
		},
	}

	for _, tc := range testCases {
		output := math.Sum[int](*tc.inputList)
		assert.Equal(t, tc.output, output)
	}
}
