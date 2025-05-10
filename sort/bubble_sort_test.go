package sort_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whitedaolee/algorithm_exercises/sort"
)

func TestBubbleSort(t *testing.T) {
	input := []int{1, 3, 2, 4, 8, 6, 8, 6, 10, 9}
	want := []int{1, 2, 3, 4, 6, 6, 8, 8, 9, 10}
	get := sort.BubbleSort(input)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want))

	input = []int{3}
	want = []int{3}
	get = sort.BubbleSort(input)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want))

	input = []int{}
	want = []int{}
	get = sort.BubbleSort(input)
	t.Log(get)
	assert.Equal(t, true, slices.Equal(get, want))
}
