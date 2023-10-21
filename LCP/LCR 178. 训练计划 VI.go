/**
 * Author: Deean
 * Date: 2023-10-19 23:34
 * FileName: LCR/LCR 178. 训练计划 VI.go
 * Description:
 */

package main

import "fmt"

func trainingPlan6(actions []int) int {
	hash := map[int]int{}
	for _, action := range actions {
		hash[action]++
	}
	for k, v := range hash {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(trainingPlan6([]int{12, 1, 6, 12, 6, 12, 6}))
}
