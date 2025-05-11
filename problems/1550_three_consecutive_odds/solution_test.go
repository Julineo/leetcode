package three_consecutive_odds

import "testing"

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "No three consecutive odds",
			arr:      []int{2, 6, 4, 1},
			expected: false,
		},
		{
			name:     "Has three consecutive odds",
			arr:      []int{1, 2, 34, 3, 4, 5, 7, 23, 12},
			expected: true,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: false,
		},
		{
			name:     "Single element",
			arr:      []int{1},
			expected: false,
		},
		{
			name:     "Two odd numbers",
			arr:      []int{1, 3},
			expected: false,
		},
		{
			name:     "Three odd numbers not consecutive",
			arr:      []int{1, 2, 3, 4, 5},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeConsecutiveOdds(tt.arr)
			if result != tt.expected {
				t.Errorf("threeConsecutiveOdds() = %v, want %v", result, tt.expected)
			}
		})
	}
}
