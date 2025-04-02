package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a number: ")
	if scanner.Scan() {
		scan := scanner.Text()
		input, err := strconv.Atoi(scan)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid input, please enter an integer:", err)
			return
		}
		result := fizzBuzz(input)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}

}
func fizzBuzz(n int) []string {
	res := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	return res
}
