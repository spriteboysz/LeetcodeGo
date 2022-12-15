/**
 * Author: Deean
 * Date: 2022/12/14 23:32
 * FileName: algorithm/P1291. 顺次数.go
 * Description:
 */

package main

import "fmt"

func sequentialDigits(low int, high int) []int {
	var target = [...]int{
		12, 23, 34, 45, 56, 67, 78, 89, 123, 234, 345, 456, 567, 678, 789,
		1234, 2345, 3456, 4567, 5678, 6789, 12345, 23456, 34567, 45678, 56789,
		123456, 234567, 345678, 456789, 1234567, 2345678, 3456789, 12345678, 23456789, 123456789,
	}
	sequential := []int{}
	for _, num := range target {
		if num >= low && num <= high {
			sequential = append(sequential, num)
		}
	}
	return sequential
}

func main() {
	fmt.Println(sequentialDigits(1000, 13000))
}
