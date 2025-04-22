// Package pkg implements algorithms and tools.
package pkg

import "fmt"

// BinarySearch implements the binary search algorithm to find a target integer in a sorted array.
func BinarySearch(arr []int, target int) (int, error) {
	if len(arr) == 0 {
		return -1, fmt.Errorf("array is empty")
	}

	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			for mid > 0 && arr[mid-1] == target {
				mid--
			}
			return mid, nil
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1, fmt.Errorf("target not found")
}
