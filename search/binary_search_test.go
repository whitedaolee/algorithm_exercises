package search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whitedaolee/algorithm_exercises/search"
)

func TestBinarySearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 6, 6, 8, 8, 12, 325, 43654, 7854321}
	target1 := 6

	get := search.BinarySearch(input, target1)
	assert.Equal(t, 5, get)

	target2 := 12
	get = search.BinarySearch(input, target2)
	assert.Equal(t, 8, get)

	target3 := 72849578924
	get = search.BinarySearch(input, target3)
	assert.Equal(t, -1, get)
}
