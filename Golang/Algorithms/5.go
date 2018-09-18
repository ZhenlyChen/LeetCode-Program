package Algorithms

import "fmt"

func Test5() {
	fmt.Println("test begin:")
	data := []string{"babad", "cbbd", "cb", "bb", ""}
	for _, t := range data {
		fmt.Println(t, "->", longestPalindrome(t))
	}
	fmt.Println("test end.")
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) < 2 {
		return s[0:1]
	}
	ans := s[0:1]
	maxLen := 0
	for i := 0; i < len(s); i++ {
		left, right := 1, 1
		for i-left >= 0 && s[i-left] == s[i] {
			left++
		}
		for i+right < len(s) && s[i+right] == s[i] {
			right++
		}
		for i-left >= 0 && i+right < len(s) && s[i-left] == s[i+right] {
			left++
			right++
		}
		if right+left > 2 && right+left-1 >= maxLen {
			maxLen = right + left - 1
			ans = s[i-left+1 : i+right]
		}
		// fmt.Println("debug:", i, ans)
	}
	return ans
}
