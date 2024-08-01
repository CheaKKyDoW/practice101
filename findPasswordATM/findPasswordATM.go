package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

// Test Case
// ตัวอย่างที่1

// Input:HE45L32LO458T6H359ISIS1BO589RNT34ODEVN80AJA
// Output:1604

// ตัวอย่างที่2
// Input:NPoeJ43OP-*!@#8805j3n62df0
// Output:8913

// ตัวอย่างที่3
// Input:a1b2c3d4e5f6g7h8i9
// Output:0045

// Good version
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line of input
	scanner.Scan()
	input := scanner.Text()

	// Initialize a slice to hold numeric substrings
	var numericParts []string
	currentNumber := ""

	// Iterate through each character in the input string
	for _, char := range input {
		if unicode.IsNumber(char) {
			currentNumber += string(char)
		} else if currentNumber != "" {
			numericParts = append(numericParts, currentNumber)
			currentNumber = ""
		}
	}
	if currentNumber != "" {
		numericParts = append(numericParts, currentNumber)
	}

	// Calculate the sum of all numeric substrings
	sum := 0
	for _, part := range numericParts {
		num, _ := strconv.Atoi(part)
		sum += num
	}

	// Convert the sum to a string and pad with leading zeros if necessary
	sumStr := fmt.Sprintf("%04d", sum)

	fmt.Println(sumStr)
}

// Bad version
func old() {

	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line
	scanner.Scan()
	s := scanner.Text()

	str_append := []string{} //make([]int, 0)
	count_append := 0
	str_append = append(str_append, "")
	res := 0
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			str_append[count_append] = fmt.Sprintf("%s%c", str_append[count_append], s[i])
		} else if str_append[count_append] != "" {
			str_append = append(str_append, "")
			count_append++
		}
	}
	for _, v := range str_append {
		num, _ := strconv.Atoi(v)
		res = res + num
	}
	numStr := strconv.Itoa(res)
	if len(numStr) < 4 {
		for i := len(numStr); i < 4; i++ {
			numStr = "0" + numStr
		}
	}
	fmt.Println(numStr)
}
