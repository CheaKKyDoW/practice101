package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

func isSubsequence(s string, t string) (result bool) {
	if s == "" {
		return true
	}
	sCount := 0
	for j := 0; j < len(t); j++ {
		if t[j] == s[sCount] {
			sCount++
			if sCount == len(s) {
				return true
			}
		}
	}

	return false
}

func isSubsequencePro(s, t string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}
