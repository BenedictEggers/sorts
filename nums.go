// Ben Eggers
// GNU GPL'd

package sorts

// We need to do this if we want to be able to sort slices of ints
type nums []int

func (n nums) Len() int {
	return len(n)
}
func (n nums) Less(i, j int) bool {
	return n[i] < n[j]
}
func (n nums) Swap(i, j int) {
	tmp := n[i]
	n[i] = n[j]
	n[j] = tmp
}

// These functions are to enable type nums to be a heap

func (n *nums) Push(x int) {
	*n = append(*n, x)
}

func (n *nums) Pop() interface{} {
	old := *n
	length := len(old)
	x := old[length - 1]
	*n = old[0 : length - 1]
	return x
}
