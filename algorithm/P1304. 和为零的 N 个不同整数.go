/**
 * Author: Deean
 * Date: 2022-10-20 22:57
 * FileName: algorithm/P1304. 和为零的 N 个不同整数.go
 * Description:
 */

package main

import "fmt"

func sumZero(n int) []int {
	var zero []int
	for i := 1; i <= n/2; i++ {
		zero = append(zero, i)
		zero = append(zero, -i)
	}
	if n%2 == 1 {
		zero = append(zero, 0)
	}
	return zero
}

func main() {
	fmt.Println(sumZero(5))
	fmt.Println(sumZero(4))
	fmt.Println(sumZero(1))
}
