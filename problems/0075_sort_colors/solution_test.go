package sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Already sorted",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Reverse sorted",
			nums:     []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		{
			name:     "All twos",
			nums:     []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("sortColors() = %v, want %v", tt.nums, tt.expected)
			}
		})
	}
}
