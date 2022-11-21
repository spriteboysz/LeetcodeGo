/**
 * Author: Deean
 * Date: 2022/11/20 22:55
 * FileName: algorithm/P1018. 可被 5 整除的二进制前缀.go
 * Description:
 */

package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	n := 0
	div := []bool{}
	for _, num := range nums {
		n = 2*(n%5) + num
		if n%5 == 0 {
			div = append(div, true)
		} else {
			div = append(div, false)
		}
	}
	return div
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
}
