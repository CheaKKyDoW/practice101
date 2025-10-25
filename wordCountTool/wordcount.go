package wordcounttool

import (
	"strings"
	"unicode"
)

const vowels string = "aeiouAEIOU"

func countValidWords(s string) int {
	words := strings.Fields(s)
	var validCount int
	for _, v := range words {
		if len(v) < 3 {
			continue
		}
		if !isVowel(v) {
			continue
		}
		if !isAlphanumeric(v) {
			continue
		}
		if !isConsonant(v) {
			continue
		}
		validCount++
	}

	return validCount
}

func isVowel(s string) bool {
	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			return true
		}
	}
	return false
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isConsonant(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !strings.ContainsRune(vowels, r) {
			return true
		}
	}
	return false
}
