/**
 * Author: Deean
 * Date: 2022-10-18 22:51
 * FileName: algorithm/P0762. 二进制表示中质数个计算置位.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	isPrime := func(num int) bool {
		if num < 2 {
			return false
		}
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}
	cnt := 0
	for i := left; i <= right; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countPrimeSetBits(10, 15))
}
