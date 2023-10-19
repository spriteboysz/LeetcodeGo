/**
 * Author: Deean
 * Date: 2023-10-16 22:55
 * FileName: LCP/LCR 139. 训练计划 I.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func trainingPlan(actions []int) []int {
	sort.Slice(actions, func(i, j int) bool {
		return actions[i]%2 > actions[j]%2
	})
	return actions
}

func main() {
	fmt.Println(trainingPlan([]int{1, 2, 3, 4, 5}))
}
