package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scoreOfString(s string) int {

	sum := 0
	// byteStr := []byte(s)
	for i := 0; i < len(s)-1; i++ {
		// fmt.Printf("%d : %d \n", s[i], s[i+1])
		a, _ := fmt.Println("%d-%d\n", s[i], s[i+1])
		sum += a
	}
	fmt.Println(sum)
	return sum
}

func mains() {

	//Input Example
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	//Output of Input Example

	data := strings.Split(input, "\n")
	num1, _ := strconv.Atoi(data[0])
	num2, _ := strconv.Atoi(data[1])
	if num1 > num2 {
		fmt.Printf("A")
	} else if num1 < num2 {
		fmt.Printf("B")
	} else {
		fmt.Printf("AB")
	}

}

func printMiddleLine(length int) {
	if length <= 1 {
		fmt.Println("#")
		return
	}

	fmt.Printf("#%s#\n", strings.Repeat(" ", length-2))
}

func main() {

	fmt.Println(strings.Repeat("#", 5))

	return
	for i := 1; i < 3-1; i++ {
		printMiddleLine(3)
	}

	return
	scanner := bufio.NewScanner(os.Stdin)

	// Read the first line
	scanner.Scan()
	firstInput := scanner.Text()
	num1, err := strconv.Atoi(firstInput)
	if err != nil {
		fmt.Println("Error converting first input to integer:", err)
		return
	}
	if num1 == 1 {
		fmt.Println("#")
		return
	}
	for i := 1; i <= num1; i++ {
		if i == 1 {
			for i := 0; i < num1; i++ {

				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}
		if i == num1 {
			for j := 0; j < num1; j++ {

				fmt.Printf("#")
			}
			fmt.Printf("\n")
		}

		if i != 1 && i != num1 {

			for k := 0; k < num1; k++ {
				// fmt.Println("start")
				if k == 0 || k == num1-1 {
					fmt.Printf("#")

				} else {
					fmt.Printf(" ")
				}
			}

			fmt.Printf("\n")
		}
	}

}
