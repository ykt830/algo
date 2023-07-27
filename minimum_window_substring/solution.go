package minimum_window_substring

import "math"

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

注意：
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

示例 2：
输入：s = "a", t = "a"
输出："a"
解释：整个字符串 s 是最小覆盖子串。

示例 3:
输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。

提示：
m == s.length
n == t.length
1 <= m, n <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
*/

func minWindow(s string, t string) string {
	set, lack := make(map[uint8]int), make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		set[t[i]]++
		lack[t[i]]++
	}
	lackCnt := len(lack)

	l, r := 0, 0
	retL, retR := 0, math.MaxInt
	for l <= len(s)-len(t)+1 {
		for ; r < len(s); r++ {
			rCh := s[r]
			if _, ok := set[rCh]; !ok {
				continue
			}

			lack[rCh]--
			if lack[rCh] == 0 {
				lackCnt--
				if lackCnt == 0 {
					break
				}
			}
		}
		if lackCnt != 0 {
			break
		}

		for r-l+1 >= len(t) {
			if r-l < retR-retL {
				retL, retR = l, r
			}

			lCh := s[l]
			l++
			if _, ok := set[lCh]; !ok {
				continue
			}
			lack[lCh]++
			if lack[lCh] == 1 {
				lackCnt++
				r++
				break
			}
		}
	}

	if retR-retL > len(s) {
		return ""
	}

	return s[retL : retR+1]
}
