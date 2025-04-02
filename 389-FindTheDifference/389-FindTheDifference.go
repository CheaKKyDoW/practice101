package main

import "fmt"

func main() {
	s, t := "abcd", "abcde"
	// s, t := "", "y"
	r := findTheDifference(s, t)
	fmt.Printf("%c", r)
}

// Pro version
func findTheDifference(s string, t string) byte {
	var result byte
	for i := range s {
		result ^= s[i] ^ t[i]
	}
	return result ^ t[len(t)-1]
}

// Bad version
// func findTheDifference(s string, t string) byte {
// 	if len(s) == 0 {
// 		return []byte(t)[0]
// 	} else if len(t) > len(s) {
// 		return t[len(s)]
// 	}
// 	result := byte(0)
// 	for i := 0; i < len(t); i++ {
// 		if s[i]^t[i] != 0 {
// 			return t[i]
// 		}
// 	}

// 	return result
// }
