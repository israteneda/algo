package sorting

func SelectionSort(elements []int) []int {
	for i := 0; i < len(elements); i++ {
		minIndex := i
		for j := i; j < len(elements); j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		elements[i], elements[minIndex] = elements[minIndex], elements[i]
	}
	return elements
}
