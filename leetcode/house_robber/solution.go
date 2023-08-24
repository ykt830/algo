package house_robber

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
给定一个代表每个房屋存放金额的非负整数数组，计算你不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

示例 1：
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。

示例 2：
输入：[2,7,9,3,1]
输出：12
解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 400
*/

func rob(nums []int) int {
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	if len(nums) > 1 {
		tmp[1] = nums[1]
	}
	if len(nums) > 2 {
		tmp[2] = nums[2] + nums[0]
	}

	for i := 3; i < len(nums); i++ {
		tmp[i] = nums[i] + max(tmp[i-2], tmp[i-3])
	}

	if len(tmp) <= 1 {
		return tmp[0]
	} else {
		return max(tmp[len(tmp)-1], tmp[len(tmp)-2])
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
