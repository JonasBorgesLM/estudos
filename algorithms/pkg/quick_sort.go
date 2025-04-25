package pkg

// QuickSort sorts an array of integers using the quicksort algorithm.
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	quickSort(arr, 0, len(arr)-1)
	return arr
}

// quickSort recursively sorts the array between the low and high indices.
func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

// partition rearranges the array around the pivot and returns its new index.
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
