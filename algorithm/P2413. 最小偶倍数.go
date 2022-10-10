/**
 * Author: Deean
 * Date: 2022-10-10 23:46
 * FileName: algorithm/P2413. 最小偶倍数.go
 * Description:
 */

package main

import "fmt"

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return n * 2
	}
}

func main() {
	fmt.Println(smallestEvenMultiple(5))
	fmt.Println(smallestEvenMultiple(6))
}
