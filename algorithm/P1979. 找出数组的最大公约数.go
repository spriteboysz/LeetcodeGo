/**
 * Author: Deean
 * Date: 2022-10-16 23:45
 * FileName: algorithm/P1979. 找出数组的最大公约数.go
 * Description:
 */

package main

import "fmt"

func findGCD(nums []int) int {
	maximum, minimum := 0, 10001
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	for i := minimum; i >= 1; i-- {
		if minimum%i == 0 && maximum%i == 0 {
			return i
		}
	}
	return 1
}

func main() {
	nums := []int{7, 5, 6, 8, 3}
	fmt.Println(findGCD(nums))
}
