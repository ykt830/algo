package container_with_most_water

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1


提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func maxArea(height []int) int {
	l, r, s := 0, len(height)-1, 0
	for l < r {
		v := (r - l) * min(height[l], height[r])
		s = max(v, s)
		if height[l] <= height[r] {
			oldL := l
			l++
			for ; l < r && height[l] <= height[oldL]; l++ {
			}
		} else {
			oldR := r
			r--
			for ; l < r && height[r] <= height[oldR]; r-- {
			}
		}
	}

	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
