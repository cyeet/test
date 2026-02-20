// Package mergesort implements the merge sort algorithm for integer slices.
package mergesort

// MergeSort returns a sorted copy of the input slice using merge sort.
// The original slice is not modified.
func MergeSort(input []int) []int {
	n := len(input)
	if n <= 1 {
		out := make([]int, n)
		copy(out, input)
		return out
	}

	half := n / 2
	sortedLeft := MergeSort(input[:half])
	sortedRight := MergeSort(input[half:])
	return mergeSorted(sortedLeft, sortedRight)
}

// mergeSorted combines two sorted slices into a single sorted slice.
func mergeSorted(a, b []int) []int {
	combined := make([]int, 0, len(a)+len(b))
	ai, bi := 0, 0
	for ai < len(a) && bi < len(b) {
		if a[ai] <= b[bi] {
			combined = append(combined, a[ai])
			ai++
		} else {
			combined = append(combined, b[bi])
			bi++
		}
	}
	combined = append(combined, a[ai:]...)
	combined = append(combined, b[bi:]...)
	return combined
}
