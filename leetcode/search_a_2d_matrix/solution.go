package search_a_2d_matrix

/*
给你一个满足下述两条属性的 m x n 整数矩阵：
每行中的整数从左到右按非递减顺序排列。
每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 100
-104 <= matrix[i][j], target <= 104
*/

func searchMatrix(matrix [][]int, target int) bool {
	index := 0
	for ; index < len(matrix); index++ {
		if matrix[index][0] > target {
			index--
			break
		}
	}
	if index == -1 {
		return false
	}
	if index == len(matrix) {
		index--
	}

	return binarySearch(matrix[index], target)
}

func binarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		m := (l + r) / 2
		if arr[m] < target {
			l = m + 1
		} else if arr[m] == target {
			return true
		} else {
			r = m - 1
		}
	}

	return false
}
