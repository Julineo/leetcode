package two_sum

// twoSum finds two numbers in the array that add up to the target.
// Returns the indices of the two numbers.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}
