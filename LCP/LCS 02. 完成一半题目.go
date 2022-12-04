/**
 * Author: Deean
 * Date: 2022/12/4 17:25
 * FileName: LCP/LCS 02. 完成一半题目.go
 * Description:
 */

package main

import "fmt"

func halfQuestions(questions []int) int {
	types := [1001]int{}
	for _, question := range questions {
		types[question]++
	}
	for i := 0; i < len(types); i++ {
		for j := i + 1; j < len(types); j++ {
			if types[i] < types[j] {
				types[i], types[j] = types[j], types[i]
			}
		}
	}

	cnt := 0
	for i := 0; i <= len(questions)/2; i++ {
		cnt += types[i]
		if cnt >= len(questions)/2 {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(halfQuestions([]int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6}))
}
