package pkg

// SelectionSort sorts an array using the selection sort algorithm.
func SelectionSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)

	n := len(sortedArr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if sortedArr[j] < sortedArr[minIndex] {
				minIndex = j
			}
		}

		sortedArr[i], sortedArr[minIndex] = sortedArr[minIndex], sortedArr[i]
	}
	return sortedArr
}
