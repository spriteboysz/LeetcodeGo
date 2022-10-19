/**
 * Author: Deean
 * Date: 2022-10-19 23:32
 * FileName: algorithm/P1619. 删除某些元素后的数组均值.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n, sum := len(arr), 0
	for _, num := range arr[n/20 : 19*n/20] {
		sum += num
	}
	return float64(sum*10) / float64(n*9)
}

func main() {
	arr := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	fmt.Println(trimMean(arr))
}
