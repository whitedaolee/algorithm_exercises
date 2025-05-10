package sort_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whitedaolee/algorithm_exercises/sort"
)

var input = []int{1, 3, 2, 4, 8, 6, 8, 6, 10, 9}
var want = []int{1, 2, 3, 4, 6, 6, 8, 8, 9, 10}

var input_one = []int{3}
var want_one = []int{3}

var input_zero = []int{}
var want_zero = []int{}

func TestBubbleSort(t *testing.T) {
	get := sort.BubbleSort(input)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want))

	get = sort.BubbleSort(input_one)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want_one))

	get = sort.BubbleSort(input_zero)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want_zero))
}

func TestSelectionSort(t *testing.T) {
	get := sort.SelectionSort(input)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want))

	get = sort.SelectionSort(input_one)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want_one))

	get = sort.SelectionSort(input_zero)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want_zero))
}
