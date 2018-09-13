package Algorithms

import "fmt"

func Test5() {
	ans1 := longestPalindrome("babad")
	ans2 := longestPalindrome("cbbd")
	ans3 := longestPalindrome("cb")
	ans4 := longestPalindrome("bb")
	ans5 := longestPalindrome("")
	fmt.Println(ans1, ans2, ans3, ans4, ans5)
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < 3 {
		return s[0:1]
	}
	ans := s[0:1]

	maxLen := 0
	for i := 0; i < len(s); i++ {
		left, right := 1, 1
		for i - left >= 0 {
			if s[i - left] == s[i] {
				left++
				continue
			} else if i + right < len(s) && s[i - left] == s[i + right] {
				right++
				continue
			}
		}
		if right + left > 2  && right + left - 1 >= maxLen{
			maxLen = right + left - 1
			ans = s[i - left + 1: i + right]
		}
	}
	return ans
}
