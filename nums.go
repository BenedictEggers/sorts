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