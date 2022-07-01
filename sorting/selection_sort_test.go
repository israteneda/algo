package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	unsorted_list := []int{5, 3, 2, 1, 4}
	expected_list := []int{1, 2, 3, 4, 5}

	sorted_list := SelectionSort(unsorted_list)

	assert.Equal(t, expected_list, sorted_list, "they should be equal")
}
