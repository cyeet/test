// Package sort provides sorting algorithms.
package sort

// Ordered is a constraint that permits any ordered type: any type that supports
// the operators < <= >= >.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Merge sorts a slice of any ordered type using the merge sort algorithm and
// returns the sorted result in a new slice.
//
// Merge sort is a stable, comparison-based, divide-and-conquer algorithm.
//
//   - Time complexity:  O(n log n) in all cases (best, average, worst).
//   - Space complexity: O(n) auxiliary space for the temporary buffer.
//   - Stability:        stable – equal elements preserve their original order.
//
// The implementation reuses a single auxiliary buffer to avoid repeated
// allocations during the merge phase.
func Merge[T Ordered](data []T) []T {
	n := len(data)
	if n <= 1 {
		out := make([]T, n)
		copy(out, data)
		return out
	}

	// Work on a copy so the caller's slice is not modified.
	out := make([]T, n)
	copy(out, data)

	// Allocate one auxiliary buffer for the entire sort.
	buf := make([]T, n)

	mergeSort(out, buf)
	return out
}

// mergeSort recursively splits and merges src in-place, using buf as scratch space.
func mergeSort[T Ordered](src, buf []T) {
	n := len(src)
	if n <= 1 {
		return
	}

	mid := n / 2
	mergeSort(src[:mid], buf[:mid])
	mergeSort(src[mid:], buf[mid:])
	mergeParts(src, mid, buf)
}

// mergeParts merges the two sorted halves src[:mid] and src[mid:] back into src,
// using buf as a temporary buffer.
func mergeParts[T Ordered](src []T, mid int, buf []T) {
	copy(buf, src)

	left := buf[:mid]
	right := buf[mid:]
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			src[k] = left[i]
			i++
		} else {
			src[k] = right[j]
			j++
		}
		k++
	}

	// Copy any remaining elements from either half.
	for i < len(left) {
		src[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		src[k] = right[j]
		j++
		k++
	}
}
