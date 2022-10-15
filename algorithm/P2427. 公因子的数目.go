/**
 * Author: Deean
 * Date: 2022-10-15 23:07
 * FileName: algorithm/P2427. 公因子的数目.go
 * Description:
 */

package main

import (
	"fmt"
)

func commonFactors(a int, b int) int {
	cnt := 0
	min := func(a int, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(commonFactors(25, 30))
}
