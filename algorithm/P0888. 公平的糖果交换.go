/**
 * Author: Deean
 * Date: 2022/11/13 16:59
 * FileName: algorithm/P0888. 公平的糖果交换.go
 * Description:
 */

package main

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sum := func(nums []int) (s int) {
		for _, num := range nums {
			s += num
		}
		return
	}

	has := func(target int, nums []int) bool {
		for _, num := range nums {
			if num == target {
				return true
			}
		}
		return false
	}

	ax, by := sum(aliceSizes), sum(bobSizes)
	ret := []int{}
	for _, v := range aliceSizes {
		y := v + (by-ax)/2
		if has(y, bobSizes) {
			ret = append(ret, v, y)
			return ret
		}
	}
	return ret
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 2, 5}, []int{2, 4}))
}
