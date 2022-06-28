package sorting

func InsertionSort(elements []int) []int {
  for index := 0; index < len(elements); index++ {
    current := index
    for current > 0 && elements[current] < elements[current - 1] {
      currentValue := elements[current]
      elements[current] = elements[current - 1]
      elements[current - 1] = currentValue
      current -= 1
    }
  }
  return elements
}
