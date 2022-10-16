/**
 * Author: Deean
 * Date: 2022-10-16 16:29
 * FileName: algorithm/P1295. 统计位数为偶数的数字.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	cnt := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	nums := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(nums))
}
