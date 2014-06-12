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

var sortFuncs = []sortFunc{
	{"bubblesort", bubbleSort},
	{"selection sort", selectionSort}}

func TestSorts2Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 2, t)
	}
}

func TestSorts3Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 3, t)
	}
}

func TestSorts10Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 10, t)
	}
}

func TestSorts50Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 50, t)
	}
}

func TestSorts100Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 100, t)
	}
}

func TestSorts500Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 500, t)
	}
}

func TestSorts1000Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 1000, t)
	}
}

func TestSorts5000Elements(t *testing.T) {
	for _, s := range sortFuncs {
		runThreeTests(s, 5000, t)
	}
}

// Helper functions

func runThreeTests(s sortFunc, n int, t *testing.T) {
	ns := getDecreasingSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name+" failed to sort", n, "decreasing elements")
	}

	ns = getIncreasingSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name+" failed to sort", n, "increasing elements")
	}

	ns = getRandomSlice(n)
	s.fn(ns)
	if !sort.IsSorted(ns) {
		t.Error(s.name+" failed to sort", n, "random elements")
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
