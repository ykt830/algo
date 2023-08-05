package permutations

/*
给定一个不含重复数字的数组 nums ，返回其所有可能的全排列。你可以按任意顺序返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

提示：
1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	ret := make([][]int, 0)
	for i, v := range nums {
		newNums, leftNums, rightNums := make([]int, 0), nums[:i], nums[i+1:]
		newNums = append(newNums, leftNums...)
		newNums = append(newNums, rightNums...)

		children := permute(newNums)

		for _, child := range children {
			newChild := append(child, v)
			ret = append(ret, newChild)
		}
	}

	return ret
}
