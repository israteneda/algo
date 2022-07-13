package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	unsortedArray := []int{6, 8, 7, 5, 9}
	expectedArray := []int{5, 6, 7, 8, 9}

	sortedArray := BubbleSort(unsortedArray)
	assert.Equal(t, expectedArray, sortedArray, "they should be equal")
}
