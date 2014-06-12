// Ben Eggers
// GNU GPL'd

package sorts

import "sort"

// Go's standard library's stable sort
func stdStableSort(data sort.Interface) {
	sort.Stable(data)
}
