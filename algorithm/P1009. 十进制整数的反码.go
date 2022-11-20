/**
 * Author: Deean
 * Date: 2022/11/20 16:57
 * FileName: algorithm/P1009. 十进制整数的反码.go
 * Description:
 */

package main

import "fmt"

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	high := 0
	for i := 1; i < 31; i++ {
		if n < (1 << i) {
			break
		}
		high = i
	}
	return n ^ ((1 << (high + 1)) - 1)
}

func main() {
	fmt.Println(bitwiseComplement(5))
}
