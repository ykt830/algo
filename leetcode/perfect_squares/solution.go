package perfect_squares

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。
完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

示例 1：
输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

示例 2：
输入：n = 13
输出：2
解释：13 = 4 + 9

提示：
1 <= n <= 104
*/

var rets [10001]int

func init() {
	squareNums := genSquareNums()
	for _, num := range squareNums {
		rets[num] = 1
	}

	for i := 1; i < len(rets); i++ {
		for j := 0; j < len(squareNums); j++ {
			v := i + squareNums[j]
			if v > 10000 {
				break
			}

			if rets[v] == 0 {
				rets[v] = rets[i] + 1
			} else {
				rets[v] = min(rets[i]+1, rets[v])
			}
		}
	}
}

func genSquareNums() []int {
	ret := make([]int, 0)
	for i := 1; i <= 100; i++ {
		ret = append(ret, i*i)
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	return rets[n]
}
