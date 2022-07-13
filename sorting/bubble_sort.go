package sorting

func BubbleSort(array []int) []int {
	l := len(array) - 1
	for i := l; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			return array
		}
	}
	return array
}
