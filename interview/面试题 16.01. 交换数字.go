/**
 * Author: Deean
 * Date: 2022-10-16 00:00
 * FileName: interview/面试题 16.01. 交换数字.go
 * Description:
 */

package main

import "fmt"

func swapNumbers(numbers []int) []int {
	a, b := numbers[0], numbers[1]
	return []int{b, a}
}

func main() {
	numbers := []int{1, 2}
	fmt.Println(swapNumbers(numbers))
}
