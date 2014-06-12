// Ben Eggers
// GNU GPL'd

package sorts

// Bubble sort

import (
	"sort"
)

func bubbleSort(data sort.Interface) {
	keepGoing := true
	for !keepGoing {
		keepGoing = false
		for i := 0; i < data.Len() - 1; i++ {
			if data.Less(i, i+1) {
				data.Swap(i, i+1)
				keepGoing = true
			}
		}
	}
}