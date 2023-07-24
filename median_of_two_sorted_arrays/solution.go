package median_of_two_sorted_arrays

import "math"

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tv1, tv2 float64

	a, b, step, length := -1, -1, 0, len(nums1)+len(nums2)
	for a+1 < len(nums1) || b+1 < len(nums2) {
		av, bv := math.MaxInt, math.MaxInt
		if a+1 < len(nums1) {
			av = nums1[a+1]
		}
		if b+1 < len(nums2) {
			bv = nums2[b+1]
		}

		var v int
		if av < bv {
			a++
			v = nums1[a]
		} else {
			b++
			v = nums2[b]
		}
		step++

		if step == length/2 {
			tv1 = float64(v)
		} else if step == length/2+1 {
			tv2 = float64(v)
		}

		if step == length/2+1 {
			if length%2 == 0 {
				return (tv1 + tv2) / 2
			} else {
				return tv2
			}
		}
	}

	return 0
}
