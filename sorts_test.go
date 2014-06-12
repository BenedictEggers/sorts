// Ben Eggers
// GNU GPL'd

// Tests the sorts

package sorts

import (
	"math/rand"
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
	for _, s := range sortFuncs {
		runThreeTests(s, 2, t)
	}
}

func TestSortsThreeElements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 3, t)
	}
}

func TestSortsTenElements(t *testing.T) {

}

// Helper functions

func runThreeTests(s sortFunc, n int, t *testing.T) {
	ns := getDecreasingSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name + " failed to sort " + string(n) + " decreasing elements")
	}

	ns = getIncreasingSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name + " failed to sort " + string(n) + " decreasing elements")
	}

	ns = getRandomSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name + " failed to sort " + string(n) + " decreasing elements")
	}
}

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
		retVal[n-i-1] = i
	}
	return retVal
}

func getRandomSlice(n int) nums {
	rand.Seed(42)
	return rand.Perm(n)
}
