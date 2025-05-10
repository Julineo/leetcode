package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Same numbers",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Same number twice",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "No solution",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum() = %v, want %v", result, tt.expected)
			}
		})
	}
}
