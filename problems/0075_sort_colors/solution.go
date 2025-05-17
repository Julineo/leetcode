package sort_colors

// sortColors sorts an array of integers where each integer is 0, 1, or 2.
// Uses the Dutch National Flag algorithm (three-way partitioning).
func sortColors(nums []int) {
	p0, curr := 0, 0
	p2 := len(nums) - 1
	for curr <= p2 {
		if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		} else if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr]
			p0++
			curr++
		} else {
			curr++
		}
	}
}
