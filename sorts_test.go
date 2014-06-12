// Ben Eggers
// GNU GPL'd

// Tests the sorts

package sorts

import (
	"sort"
	"testing"
	"math/rand"
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
	for _, s := range sortFuncs {
		n := nums{2, 1}
		s.fn(n)
		if !sort.IsSorted(n) {
			// Didn't swap the elements
			t.Error(s.name + " failed to sort 2 elements")
		}
	}
}

func TestSortsThreeElements(t *testing.T) {
	for _, s := range sortFuncs {
		n := nums{2, 1, 5}
		s.fn(n)
		if !sort.IsSorted(n) {
			// Didn't swap the elements
			t.Error(s.name + " failed to sort 3 elements")
		}
	}
}

func TestSortsTenElements(t *testing.T) {

}




// Helper functions

func getIncreasingSlice(n int) nums {
	retVal := make(nums, n)
	for i := 0; i < n; i++ {
		retVal[i] = i
	}
	return retVal
}

func getDecreasingSlice(n int) nums {
	retVal := make(nums, n)
	for i := 0; i < n; i++ {
		retVal[n - i - 1] = i
	}
	return retVal
}

func getRandomSlice(n int) nums {
	// Randomness is bad in tests, but whatever
	return rand.Perm(n)
}