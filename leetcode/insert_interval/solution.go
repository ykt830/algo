package insert_interval

/*
给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]

示例 2：
输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

示例 3：
输入：intervals = [], newInterval = [5,7]
输出：[[5,7]]

示例 4：
输入：intervals = [[1,5]], newInterval = [2,3]
输出：[[1,5]]

示例 5：
输入：intervals = [[1,5]], newInterval = [2,7]
输出：[[1,7]]

提示：
0 <= intervals.length <= 10^4
intervals[i].length == 2
0 <= intervals[i][0] <= intervals[i][1] <= 10^5
intervals 根据 intervals[i][0] 按 升序 排列
newInterval.length == 2
0 <= newInterval[0] <= newInterval[1] <= 10^5
*/

var ret = make([][]int, 0)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	isSet := false
	ret = append([][]int{}, []int{-1, -1})
	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		if !isSet && cur[0] > newInterval[0] {
			cur = newInterval
			isSet = true
			i--
		}

		add(cur)
	}

	if !isSet {
		add(newInterval)
	}

	return ret[1:]
}

func add(v []int) {
	top := ret[len(ret)-1]
	if v[0] <= top[1] && v[1] >= top[1] {
		top[1] = v[1]
	} else if v[0] > top[1] {
		ret = append(ret, v)
	}
}
