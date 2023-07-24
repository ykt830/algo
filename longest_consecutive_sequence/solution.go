package longest_consecutive_sequence

import "math"

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9


提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := map[int]int{}
	for _, v := range nums {
		m[v] = 1
	}

	max := 1
	for k, v := range m {
		if v == 0 {
			continue
		}

		count := 1
		for i := k + 1; m[i] == 1; i++ {
			count++
			m[i] = 0
		}
		for i := k - 1; m[i] == 1; i-- {
			count++
			m[i] = 0
		}

		max = int(math.Max(float64(max), float64(count)))
	}

	return max
}
