package jump_game

func canJump2(nums []int) bool {
	maxP := nums[0]
	for i := 1; i < len(nums) && i <= maxP; i++ {
		newP := i + nums[i]
		if newP > maxP {
			maxP = newP
			if maxP >= len(nums)-1 {
				return true
			}
		}
	}

	return maxP >= len(nums)-1
}
