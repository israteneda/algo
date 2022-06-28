package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	unsortedElements := []int{4, 5, 3, 2, 7}
	expectedElements := []int{2, 3, 4, 5, 7}

	sortedElements := InsertionSort(unsortedElements)

	assert.Equal(t, sortedElements, expectedElements, "they should be equal")
}
