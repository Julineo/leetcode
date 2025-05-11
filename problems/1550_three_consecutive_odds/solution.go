package three_consecutive_odds

// threeConsecutiveOdds checks if there are three consecutive odd numbers in the array.
func threeConsecutiveOdds(arr []int) bool {
	oddsCount := 0
	for _, num := range arr {
		if num&1 == 1 { // bitwise num%2 != 0
			oddsCount++
			if oddsCount == 3 {
				return true
			}
		} else {
			oddsCount = 0
		}
	}
	return false
}
