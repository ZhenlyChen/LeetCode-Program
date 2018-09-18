package Algorithms

import "fmt"

func Test10() {
	fmt.Println("test begin:")
	s := []string{"aa", "aa", "ab", "aab", "mississippi", "ab", "aaa", ""}
	p := []string{"a", "a*", ".*", "c*a*b", "mis*is*p*", ".*c", "a*a", "c*c*"}
	ans := []bool{false, true, true, true, false, false, true, true}
	for i := range s {
		fmt.Println("s:", s[i], "p:", p[i], "->", isMatch(s[i], p[i]), "ans:", ans[i])
	}
	fmt.Println("test end.")
}

// 以下三个函数其实方法都一样，只是缩短了一下而已


func isMatch3(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	if sLen <= 0 || pLen <= 0 {
		if sLen <= 0 && pLen <= 0 {
			return true
		} else if pLen >= 2 && p[1] == '*' {
			return isMatch3(s[:], p[2:])
		} else {
			return false
		}
	}

	if p[0] >= 'a' && p[0] <= 'z' {
		if pLen > 1 && p[1] == '*' {
			return isMatch3(s[:], p[2:]) || (p[0] == s[0] && isMatch3(s[1:], p[:]))
		} else {
			return p[0] == s[0] && isMatch3(s[1:], p[1:])
		}
	} else if p[0] == '.' {
		if pLen > 1 && p[1] == '*' {
			return isMatch3(s[:], p[2:]) || isMatch3(s[1:], p[:])
		} else {
			return isMatch3(s[1:], p[1:])
		}
	}
	return false
}

func isMatch2(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	if sLen <= 0 || pLen <= 0 {
		return sLen <= 0 && pLen <= 0 || (pLen >= 2 && p[1] == '*' && isMatch2(s[:], p[2:]))
	}
	match := p[0] == '.' || p[0] == s[0]
	if pLen > 1 && p[1] == '*' {
		return isMatch2(s[:], p[2:]) || (match && isMatch2(s[1:], p[:]))
	} else {
		return match && isMatch2(s[1:], p[1:])
	}
}


func isMatch(s string, p string) bool {
	return ((len(s) <= 0 || len(p) <= 0) && (len(s) <= 0 && len(p) <= 0 || (len(p) >= 2 && p[1] == '*' && isMatch(s[:], p[2:])))) || (!(len(s) <= 0 || len(p) <= 0)&&((len(p) > 1 && p[1] == '*' && (isMatch(s[:], p[2:]) || ((p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[:])))) ||  ((p[0] == '.' || p[0] == s[0]) && isMatch(s[1:], p[1:]))))
}