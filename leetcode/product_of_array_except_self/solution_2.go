package product_of_array_except_self

func productExceptSelf2(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1

	for i := 1; i < len(nums); i++ {
		ret[i] = nums[i-1] * ret[i-1]
	}

	v := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if i == 0 {
			ret[i] = v
			continue
		}

		ret[i] *= v
		v *= nums[i]
	}

	return ret
}
