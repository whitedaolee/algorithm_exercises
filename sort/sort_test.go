package sort_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whitedaolee/algorithm_exercises/sort"
)

var inputs = [][]int{
	{1, 3, 2, 4, 8, 6, 8, 6, 10, 9},
	{3},
	{},
}

var wants = [][]int{
	{1, 2, 3, 4, 6, 6, 8, 8, 9, 10},
	{3},
	{},
}

func sortTest(t *testing.T, sortFunc func([]int) []int) {
	for i, input := range inputs {
		get := sortFunc(input)
		t.Log(get)
		assert.Equal(t, true, slices.Equal(get, wants[i]))
	}
}

func TestBubbleSort(t *testing.T) {
	sortTest(t, sort.BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	sortTest(t, sort.SelectionSort)
}

func TestInsertionSort(t *testing.T) {
	sortTest(t, sort.InsertionSort)
}
