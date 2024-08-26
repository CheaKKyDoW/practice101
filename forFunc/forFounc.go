package main

import (
	"fmt"
	"iter"
	"strings"
	"unicode/utf8"
)

func WordCount(s string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, v := range strings.Split(s, " ") {
			if !yield(utf8.RuneCountInString(v)) {
			}
		}
	}
}

func main() {
	s := "my life is more difficult, อยากจะร้องง"
	for v := range WordCount(s) {
		fmt.Println(v, " ")
	}
}
