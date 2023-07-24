package group_anagrams

import "sort"

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for i := range strs {
		arr := m[sortStr(strs[i])]
		if arr == nil {
			arr = make([]string, 0)
		}
		arr = append(arr, strs[i])
		m[sortStr(strs[i])] = arr
	}

	ret := make([][]string, 0)
	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

func sortStr(s string) string {
	if len(s) <= 1 {
		return s
	}

	letters := make([]rune, 0, len(s))
	for _, c := range s {
		letters = append(letters, c)
	}
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	var ret string
	for i := range letters {
		ret += string(letters[i])
	}

	return ret
}
