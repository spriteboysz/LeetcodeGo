/**
 * Author: Deean
 * Date: 2023-10-19 22:55
 * FileName: LCR/LCR 172. 统计目标成绩的出现次数.go
 * Description:
 */

package main

import "fmt"

func countTarget(scores []int, target int) int {
	cnt := 0
	for _, score := range scores {
		if score == target {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countTarget([]int{2, 2, 3, 4, 4, 4, 5, 6, 6, 8}, 4))
}
