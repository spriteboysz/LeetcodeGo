/**
 * Author: Deean
 * Date: 2022-11-08 22:51
 * FileName: algorithm/P1356. 根据数字二进制下 1 的数目排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	process := func(num int) (count int) {
		for ; num > 0; num /= 2 {
			count += num % 2
		}
		return
	}
	sort.Slice(arr, func(i, j int) bool {
		x, y := arr[i], arr[j]
		cx, cy := process(x), process(y)
		return cx < cy || cx == cy && x < y
	})
	return arr
}

func main() {
	fmt.Println(sortByBits([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
}
