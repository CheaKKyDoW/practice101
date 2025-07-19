package main

import "fmt"

func main() {
	var (
		nums   = []int{2, 7, 6, 11, 15}
		target = 9
	)
	result := findTwoSum(nums, target)

	fmt.Println("findTwoSum :", result)

}

func findTwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num // 9 - 2 = 7
		if iNumMap, found := numMap[complement]; found {
			return []int{iNumMap, i}
		}
		numMap[num] = i
	}
	return nil
}

// Bad Version
// func findTwoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nil
// }
