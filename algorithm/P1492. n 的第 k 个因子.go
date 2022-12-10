/**
 * Author: Deean
 * Date: 2022/12/10 17:03
 * FileName: algorithm/P1492. n 的第 k 个因子.go
 * Description:
 */

package main

import "fmt"

func kthFactor(n int, k int) int {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			k--
		}
		if k == 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(kthFactor(12, 3))
	fmt.Println(kthFactor(4, 4))
}
