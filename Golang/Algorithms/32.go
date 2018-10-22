package Algorithms

import "fmt"

func Test32 () {
	testData := []string{"(()", ")()())", "(()(()()(", "", "()"}
	for i := range testData {
		fmt.Println(testData[i], longestValidParentheses(testData[i]))
		fmt.Println(testData[i], longestValidParenthesesDP(testData[i]))
	}
}

func longestValidParentheses(s string) int {
	var stack []int32
	var pos []int
	for i, c := range s {
		if len(stack) != 0 && c == ')' && stack[len(stack) - 1] == '(' {
			stack = stack[:len(stack) - 1]
			pos = pos[:len(pos) - 1]
		} else {
			stack = append(stack, c)
			pos = append(pos, i)
		}
	}
	pos = append([]int{-1}, pos[:]...)
	pos = append(pos, len(s))
	// fmt.Println(pos)
	max := 0
	for i := 1; i < len(pos); i ++ {
		lengths := pos[i] - pos[i  - 1] - 1
		if lengths > max {
			max = lengths
		}
	}
	return max
}

func longestValidParenthesesDP(s string) int {
	left, right, max := 0, 0, 0
	for i := range s {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if 2 * right > max {
				max = 2 * right
			}
		} else if right >= left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if 2 * left > max {
				max = 2 * left
			}
		} else if  left >= right {
			left, right = 0, 0
		}
	}
	return max
}