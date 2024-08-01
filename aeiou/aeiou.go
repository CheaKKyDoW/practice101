package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Test Case
// ตัวอย่างที่1
// Input : Welcome to BornToDev
// Output : Wlcm t BrnTDv

// ตัวอย่างที่2
// Input : I am A HERO!!
// Output:  m  HR!!

// ตัวอย่างที่3
// Input : The quick brown fox jumps over the lazy dog.
// Output : Th qck brwn fx jmps vr th lzy dg.

// Bad version
func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	res := ""
	for _, v := range input {
		res_chk, _ := regexp.MatchString(`[AEIOUaeiou]`, strconv.QuoteRune(v))
		if !res_chk {
			res = res + fmt.Sprintf("%c", v)
		}
	}

	fmt.Println(res)
}
