package pascals_triangle

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

示例 2:
输入: numRows = 1
输出: [[1]]

提示:
1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	ret := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		before := ret[i-1]
		cur := make([]int, i+1)
		cur[0], cur[len(cur)-1] = 1, 1
		for j := 1; j < len(cur)-1; j++ {
			cur[j] = before[j-1] + before[j]
		}
		ret = append(ret, cur)
	}

	return ret
}
