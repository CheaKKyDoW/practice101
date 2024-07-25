package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// #####
// #   #
// #   #
// #   #
// #####

func main() {

	scanner := bufio.NewScanner(os.Stdin)

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
			fmt.Println(strings.Repeat("#", num1))

		}
		if i == num1 {
			fmt.Println(strings.Repeat("#", num1))
			fmt.Printf("\n")
		}

		if i != 1 && i != num1 {
			fmt.Printf("#%s#\n", strings.Repeat(" ", num1-2))

		}
	}

}

// standard
func normal_version() {

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
