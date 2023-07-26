package first_missing_positive

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


示例 1：

输入：nums = [1,2,0]
输出：3
示例 2：

输入：nums = [3,4,-1,1]
输出：2
示例 3：

输入：nums = [7,8,9,11,12]
输出：1


提示：

1 <= nums.length <= 5 * 10^5
-2^31 <= nums[i] <= 2^31 - 1
*/

func firstMissingPositive(nums []int) int {
	records := make(map[int]int, len(nums))
	for i := 1; i < len(nums)+1; i++ {
		records[i]++
	}

	for _, num := range nums {
		if _, ok := records[num]; ok {
			records[num] = 0
		}
	}

	ret := len(nums) + 1
	for k, v := range records {
		if v == 1 {
			ret = min(ret, k)
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
