package main

import (
	"fmt"
	"sort"
)

func main() {
	// Input: nums1 = [[1,2],[2,3],[4,5]], nums2 = [[1,4],[3,2],[4,1]]
	// Output: [[1,6],[2,3],[3,2],[4,6]]

	result := mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}})
	fmt.Println("Result:", result)

}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	idMap := make(map[int]int)

	for _, pair := range nums1 {
		fmt.Println("num1", pair)
		idMap[pair[0]] += pair[1]
	}
	for _, pair := range nums2 {
		fmt.Println("num2", pair)
		idMap[pair[0]] += pair[1]
	}

	keys := make([]int, 0, len(idMap))
	for v := range idMap {
		keys = append(keys, v)
	}

	sort.Ints(keys)

	result := make([][]int, 0, len(idMap))
	for _, value := range keys {
		result = append(result, []int{value, idMap[value]})
	}

	return result
}
