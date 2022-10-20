/**
 * Author: Deean
 * Date: 2022-10-20 23:02
 * FileName: algorithm/P0258. 各位相加.go
 * Description:
 */

package main

import "fmt"

func addDigits(num int) int {
	for num >= 10 {
		cur := 0
		for ; num > 0; num /= 10 {
			cur += num % 10
		}
		num = cur
	}
	return num
}

func main() {
	fmt.Println(addDigits(38))
	fmt.Println(addDigits(0))
}
