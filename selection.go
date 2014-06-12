// Ben Eggers
// GNU GPL'd

package sorts

import "sort"

// Selection sort
func selectionSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n-1; i++ {
		mindex := i // get it??
		for j := i; j < n; j++ {
			if data.Less(j, mindex) {
				mindex = j
			}
		}
		data.Swap(i, mindex)
	}
}
