/**
 * Author: Deean
 * Date: 2022/11/19 22:10
 * FileName: sword/剑指 Offer 40. 最小的k个数.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 2))
}
