/**
 * Author: Deean
 * Date: 2022-10-22 15:57
 * FileName: algorithm/P0066. 加一.go
 * Description:
 */

package main

import "fmt"

func plusOne(digits []int) []int {
	sum, carry := []int{}, 1
	for i := len(digits) - 1; i >= 0 || carry != 0; i-- {
		if i >= 0 {
			carry += digits[i]
		}
		sum = append([]int{carry % 10}, sum...)
		carry /= 10
	}
	return sum
}

func main() {
	digits := []int{9, 9, 9, 9}
	fmt.Println(plusOne(digits))
	digits = []int{2, 0, 0, 9}
	fmt.Println(plusOne(digits))
}
