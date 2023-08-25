package coin_change

import "math"

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0

提示：
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4
*/

var records = make(map[int]int)

func coinChange(coins []int, amount int) int {
	records = make(map[int]int, len(coins)+1)
	records[0] = 0
	for _, coin := range coins {
		records[coin] = 1
	}

	return calculate(coins, amount)
}

func calculate(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if v, ok := records[amount]; ok {
		return v
	}

	rets := make([]int, 0)
	for _, coin := range coins {
		rets = append(rets, calculate(coins, amount-coin))
	}

	ret := min(rets)
	if ret == -1 {
		records[amount] = ret
	} else {
		records[amount] = ret + 1
	}

	return records[amount]
}

func min(n []int) int {
	ret := math.MaxInt
	for _, v := range n {
		if v < 0 {
			continue
		}

		if v < ret {
			ret = v
		}
	}

	if ret == math.MaxInt {
		return -1
	}

	return ret
}
