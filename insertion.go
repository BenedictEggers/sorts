// Ben Eggers
// GNU GPL'd

package sorts

import "sort"

// Insertion sort
func insertionSort(data sort.Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
