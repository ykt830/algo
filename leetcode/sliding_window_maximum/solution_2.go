package sliding_window_maximum

func maxSlidingWindow2(nums []int, k int) []int {
	stack := []int{0}
	for i := 1; i < k; i++ {
		for nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
		stack = append(stack, i)
	}

	ret := []int{nums[stack[0]]}
	for i := k; i < len(nums); i++ {
		for nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
		stack = append(stack, i)

		for stack[0] < i-k+1 {
			stack = stack[1:]
		}

		ret = append(ret, nums[stack[0]])
	}

	return ret
}
