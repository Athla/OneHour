package leetcode

func isValid(s string) bool {
	closedOpen := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	if len(s)%2 != 0 {
		return false
	}

	currStack := []rune{}

	for _, v := range s {
		if _, ok := closedOpen[v]; ok {
			currStack = append(currStack, v)
		} else if len(currStack) == 0 || closedOpen[currStack[len(currStack)-1]] != v {
			return false
		} else {
			currStack = currStack[:len(currStack)-1]
		}
	}

	return len(currStack) == 0
}
