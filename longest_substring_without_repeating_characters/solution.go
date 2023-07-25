package longest_substring_without_repeating_characters

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。

	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

提示：

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, r, m, length := 0, 1, map[uint8]bool{s[0]: true}, 1
	for len(s)-l > length {
		m = map[uint8]bool{s[l]: true}
		for r = l + 1; r < len(s) && !m[s[r]]; r++ {
			m[s[r]] = true
		}
		length = max(len(m), length)
		l++
	}

	return length
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
