package mergesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "single element",
			input: []int{42},
			want:  []int{42},
		},
		{
			name:  "already sorted",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "reverse sorted",
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "with duplicates",
			input: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			want:  []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		{
			name:  "all duplicates",
			input: []int{7, 7, 7, 7},
			want:  []int{7, 7, 7, 7},
		},
		{
			name:  "two elements",
			input: []int{9, 1},
			want:  []int{1, 9},
		},
		{
			name:  "negative numbers",
			input: []int{0, -3, 5, -1, 2},
			want:  []int{-3, -1, 0, 2, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeSort(tc.input)
			if len(got) != len(tc.want) {
				t.Fatalf("MergeSort(%v) length = %d, want %d", tc.input, len(got), len(tc.want))
			}
			for idx := range tc.want {
				if got[idx] != tc.want[idx] {
					t.Errorf("MergeSort(%v)[%d] = %d, want %d", tc.input, idx, got[idx], tc.want[idx])
				}
			}
		})
	}
}

func TestMergeSortDoesNotModifyInput(t *testing.T) {
	original := []int{5, 3, 1, 4, 2}
	snapshot := make([]int, len(original))
	copy(snapshot, original)

	MergeSort(original)

	for idx := range snapshot {
		if original[idx] != snapshot[idx] {
			t.Errorf("input was modified: original[%d] = %d, want %d", idx, original[idx], snapshot[idx])
		}
	}
}

func TestMergeSortRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 20; trial++ {
		size := rng.Intn(200)
		nums := make([]int, size)
		for k := range nums {
			nums[k] = rng.Intn(1000) - 500
		}

		got := MergeSort(nums)

		expected := make([]int, size)
		copy(expected, nums)
		sort.Ints(expected)

		if len(got) != len(expected) {
			t.Fatalf("trial %d: length mismatch: got %d, want %d", trial, len(got), len(expected))
		}
		for idx := range expected {
			if got[idx] != expected[idx] {
				t.Errorf("trial %d: got[%d] = %d, want %d", trial, idx, got[idx], expected[idx])
			}
		}
	}
}
