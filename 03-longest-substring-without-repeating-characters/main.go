package lswrc

import "strings"

func lengthOfLongestSubstring(s string) int {
	var mlen int = 0
	var outer int = 0
	var pchar, char rune
	var str string
	for outer, pchar = range s {
		str = string(pchar)
		for _, char = range s[outer+1 : len(s)] {
			if strings.Contains(str, string(char)) {
				break
			}
			str += string(char)
		}
		if len(str) > mlen {
			mlen = len(str)
		}
	}
	return mlen
}
