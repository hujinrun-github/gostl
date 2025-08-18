package sort

// HeapSort sorts the slice using heap sort algorithm.
import (
	"reflect"
	"testing"
)

type IntSlice []int

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestHeapSort(t *testing.T) {
	tests := []struct {
		input    IntSlice
		expected IntSlice
	}{
		{input: IntSlice{}, expected: IntSlice{}},
		{input: IntSlice{1}, expected: IntSlice{1}},
		{input: IntSlice{5, 3, 8, 4, 2}, expected: IntSlice{2, 3, 4, 5, 8}},
		{input: IntSlice{10, 9, 8, 7, 6, 5}, expected: IntSlice{5, 6, 7, 8, 9, 10}},
		{input: IntSlice{1, 2, 3, 4, 5}, expected: IntSlice{1, 2, 3, 4, 5}},
		{input: IntSlice{3, 3, 3}, expected: IntSlice{3, 3, 3}},
		{input: IntSlice{-1, -3, 2, 0}, expected: IntSlice{-3, -1, 0, 2}},
	}

	for _, tt := range tests {
		// arr := IntSlice{}
		// copy(arr, tt.input)
		HeapSort(tt.input)
		if !reflect.DeepEqual(tt.input, tt.expected) {
			t.Errorf("HeapSort(%v) = %v; want %v", tt.input, tt.input, tt.expected)
		}
	}
}
