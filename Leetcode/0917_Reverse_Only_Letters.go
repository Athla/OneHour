package leetcode

import "unicode"

func reverseOnlyLetters(s string) string {
	str := []rune{}

	for _, v := range s {
		str = append(str, v)
	}

	l, r := 0, len(str)-1

	for l < r {
		if !unicode.IsLetter(str[l]) {
			l++
		} else if !unicode.IsLetter(str[r]) {
			r--
		} else {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
	}

	var ans string
	for _, v := range str {
		ans += string(v)
	}

	return ans
}
