package main

import "fmt"

func main() {
	s := "BorntoDev"

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}
}

// func recursive(n int) {

// 	scanner := bufio.NewScanner(os.Stdin)

// 	// Read the first line
// 	scanner.Scan()
// 	firstInput := scanner.Text()
// 	num1, err := strconv.Atoi(firstInput)
// 	if err != nil {
// 		fmt.Println("Error converting first input to integer:", err)
// 		return
// 	}
// 	res := 0
// 	for i := 0; i < num1; i++ {
// 		if res < 0 {
// 			fmt.Println(-1)
// 			return
// 		} else if res == 0 {
// 			fmt.Println("1")
// 			return
// 		} else {
// 			return recursive(num1)
// 		}
// 	}

// }
