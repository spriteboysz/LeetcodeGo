/**
 * Author: Deean
 * Date: 2022-11-10 23:00
 * FileName: algorithm/P1387. 将整数按权重排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func getKth(lo int, hi int, k int) int {
	process := func(num int) (weight int) {
		for num > 1 {
			if num%2 == 0 {
				num /= 2
			} else {
				num = num*3 + 1
			}
			weight++
		}
		return
	}
	nums := []int{}
	for i := lo; i <= hi; i++ {
		nums = append(nums, i)
	}
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		px, py := process(x), process(y)
		if px == py {
			return x < y
		}
		return px < py
	})
	return nums[k-1]
}

func main() {
	fmt.Println(getKth(7, 11, 4))
}
