package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Test Case
// ตัวอย่างที่ 1
// Input: 3
// Output:
//   *
//  ***
// *****

// ตัวอย่างที่ 2
// Input: 5
// Output:
//     *
//    ***
//   *****
//  *******
// *********

func main() {

	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	input, _ := strconv.Atoi(inputStr)
	spaceCount := input
	for i := 1; i <= input; i++ {
		fmt.Printf("%s", strings.Repeat(" ", spaceCount-1))
		spaceCount--
		for j := 0; j < i*2-1; j++ { //Why i*2-1 because it will be 3 5 7 9 from 4 6 8 10 for start *
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
