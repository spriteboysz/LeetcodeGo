/**
 * Author: Deean
 * Date: 2022/12/11 16:02
 * FileName: algorithm/面试题 17.14. 最小K个数.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func smallestK(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	fmt.Println(smallestK([]int{1, 3, 5, 7, 2, 4, 6, 8}, 4))
}
