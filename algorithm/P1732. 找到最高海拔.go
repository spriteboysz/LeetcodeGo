/**
 * Author: Deean
 * Date: 2022-10-16 22:42
 * FileName: algorithm/P1732. 找到最高海拔.go
 * Description:
 */

package main

import "fmt"

func largestAltitude(gain []int) int {
	maximum, cur := 0, 0
	for _, g := range gain {
		cur += g
		if cur > maximum {
			maximum = cur
		}
	}
	return maximum
}

func main() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}
