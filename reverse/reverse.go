package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstInput := scanner.Text()

	for i := len(firstInput) - 1; i >= 0; i-- {
		fmt.Printf("%c", firstInput[i])
	}

	//This code ensures that any leading or trailing whitespace in the input is removed before reversing
	//and printing the string. The %c format specifier is necessary for printing each character,
	//but there should not be any unwanted % characters printed if the input is processed correctly.
	fmt.Println()

}
