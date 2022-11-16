/**
 * Author: Deean
 * Date: 2022/11/16 23:21
 * FileName: algorithm/P2455. 可被三整除的偶数的平均值.go
 * Description:
 */

package main

import "fmt"

func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, num := range nums {
		if num%6 == 0 {
			sum += num
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}

func main() {
	fmt.Println(averageValue([]int{1, 3, 6, 10, 12, 15}))
}
