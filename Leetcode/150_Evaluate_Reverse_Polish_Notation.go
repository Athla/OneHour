package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	valid := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	for _, v := range tokens {
		if !valid[v] {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		} else {
			n2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			n1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch v {
			case "+":
				stack = append(stack, n1+n2)
			case "-":
				stack = append(stack, n1-n2)
			case "*":
				stack = append(stack, n1*n2)
			case "/":
				stack = append(stack, int(n1/n2))
			}
		}
	}

	return stack[0]
}
