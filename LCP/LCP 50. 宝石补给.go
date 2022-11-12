/**
 * Author: Deean
 * Date: 2022/11/12 20:05
 * FileName: LCP/LCP 50. 宝石补给.go
 * Description:
 */

package main

import "fmt"

func giveGem(gem []int, operations [][]int) int {
	for _, operation := range operations {
		start, end := operation[0], operation[1]
		cur := gem[start] / 2
		gem[start] -= cur
		gem[end] += cur
	}
	maximum, minimum := 0, 1001
	for _, num := range gem {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	return maximum - minimum
}

func main() {
	fmt.Println(giveGem([]int{100, 0, 50, 100}, [][]int{{0, 2}, {0, 1}, {3, 0}, {3, 0}}))
}
