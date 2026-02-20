package sort_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	msort "example/hello/sort"
)

// isSorted reports whether s is in non-decreasing order.
func isSorted[T msort.Ordered](s []T) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}

// isPermutation reports whether b contains exactly the same elements as a.
func isPermutation[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	counts := make(map[T]int, len(a))
	for _, v := range a {
		counts[v]++
	}
	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}
	return true
}

func TestMerge_Empty(t *testing.T) {
	got := msort.Merge([]int{})
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

func TestMerge_Single(t *testing.T) {
	input := []int{42}
	got := msort.Merge(input)
	if len(got) != 1 || got[0] != 42 {
		t.Fatalf("expected [42], got %v", got)
	}
}

func TestMerge_AlreadySorted(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	got := msort.Merge(input)
	if !isSorted(got) {
		t.Fatalf("expected sorted slice, got %v", got)
	}
	if !isPermutation(input, got) {
		t.Fatalf("result is not a permutation of input")
	}
}

func TestMerge_ReverseSorted(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	got := msort.Merge(input)
	if !isSorted(got) {
		t.Fatalf("expected sorted slice, got %v", got)
	}
	if !isPermutation(input, got) {
		t.Fatalf("result is not a permutation of input")
	}
}

func TestMerge_Duplicates(t *testing.T) {
	input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	got := msort.Merge(input)
	if !isSorted(got) {
		t.Fatalf("expected sorted slice, got %v", got)
	}
	if !isPermutation(input, got) {
		t.Fatalf("result is not a permutation of input")
	}
}

func TestMerge_NegativeNumbers(t *testing.T) {
	input := []int{-3, -1, -4, 0, 2, -7}
	got := msort.Merge(input)
	if !isSorted(got) {
		t.Fatalf("expected sorted slice, got %v", got)
	}
	if !isPermutation(input, got) {
		t.Fatalf("result is not a permutation of input")
	}
}

func TestMerge_Random(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 20; i++ {
		n := rng.Intn(200)
		input := make([]int, n)
		for j := range input {
			input[j] = rng.Intn(1000) - 500
		}

		got := msort.Merge(input)

		if !isSorted(got) {
			t.Fatalf("run %d: result not sorted: %v", i, got)
		}
		if !isPermutation(input, got) {
			t.Fatalf("run %d: result is not a permutation of input", i)
		}
	}
}

// TestMerge_DoesNotModifyInput verifies that the original slice is unchanged.
func TestMerge_DoesNotModifyInput(t *testing.T) {
	input := []int{5, 3, 1, 4, 2}
	orig := make([]int, len(input))
	copy(orig, input)

	msort.Merge(input)

	for i, v := range input {
		if v != orig[i] {
			t.Fatalf("input was modified at index %d: want %d, got %d", i, orig[i], v)
		}
	}
}

// TestMerge_Strings checks that the generic implementation works for strings.
func TestMerge_Strings(t *testing.T) {
	input := []string{"banana", "apple", "cherry", "date"}
	got := msort.Merge(input)

	stdSorted := make([]string, len(input))
	copy(stdSorted, input)
	sort.Strings(stdSorted)

	for i := range got {
		if got[i] != stdSorted[i] {
			t.Fatalf("index %d: want %q, got %q", i, stdSorted[i], got[i])
		}
	}
}

// TestMerge_Stability verifies that equal elements maintain their relative order.
//
// Items are encoded as key*1000 + originalIndex. Within the same "key group"
// (elements that share the same key*1000 prefix), a stable sort must output
// them in ascending original-index order.
//
//	Keys (original order):  3, 1, 2, 1, 3, 2  (indices 0..5)
//	Encoded:                3000, 1001, 2002, 1003, 3004, 2005
//	Expected stable output: 1001, 1003, 2002, 2005, 3000, 3004
func TestMerge_Stability(t *testing.T) {
	input := []int{3000, 1001, 2002, 1003, 3004, 2005}
	got := msort.Merge(input)
	want := []int{1001, 1003, 2002, 2005, 3000, 3004}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("stability: index %d: want %d, got %d", i, want[i], got[i])
		}
	}
}

// Example demonstrates usage of Merge with a slice of integers.
func Example() {
	sorted := msort.Merge([]int{5, 2, 8, 1, 9, 3})
	fmt.Println(sorted)
	// Output: [1 2 3 5 8 9]
}
