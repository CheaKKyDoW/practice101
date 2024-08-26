package main

//
// func countSeniors() int {
// 	input := []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}
// 	countOldpeople := 0
// 	for _, v := range input {
// 		age, _ := strconv.Atoi(v[11:13])
// 		if age > 60 {
// 			countOldpeople++
// 		}
// 	}
// 	return countOldpeople
// }

// New refector
func countSeniors(details []string) int {
	countOldpeople := 0
	for _, v := range details {
		if v[11] > '6' || (v[11] == '6' && v[12] > '0') {
			countOldpeople++
		}
	}
	return countOldpeople
}
