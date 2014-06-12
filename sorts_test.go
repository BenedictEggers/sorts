// Ben Eggers
// GNU GPL'd

// Tests the sorts

package sorts

import (
	"sort"
	"testing"
)

// We declare this type so we can have a slice of functions (making it easy to
// run them all on each test), and we can also track which one we're
// running with the "name" field
type sortFunc struct {
	name string
	fn   func(sort.Interface)
}

var sortFuncs = []sortFunc{{"bubblesort", bubbleSort}}

func TestSortsTwoElements(t *testing.T) {
	for _, sort := range sortFuncs {
		n := nums{2, 1}
		sort.fn(n)
		if n[0] == 2 {
			// Didn't swap the elements
			t.Error(sort.name + " failed to sort 2 elements")
		}
	}
}