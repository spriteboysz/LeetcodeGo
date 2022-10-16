/**
 * Author: Deean
 * Date: 2022-10-16 17:09
 * FileName: algorithm/P1051. 高度检查器.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	var sorted []int
	for _, height := range heights {
		sorted = append(sorted, height)
	}
	sort.Ints(sorted)
	cnt := 0
	for i, height := range heights {
		if height != sorted[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	heights := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(heights))
}
