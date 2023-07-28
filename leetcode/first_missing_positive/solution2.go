package first_missing_positive

func firstMissingPositive2(nums []int) int {
	for i := range nums {
		if nums[i] > len(nums) || nums[i] < 1 {
			nums[i] = len(nums) + 1
		}
	}

	for i := range nums {
		if abs(nums[i]) != len(nums)+1 {
			nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
		}
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}
