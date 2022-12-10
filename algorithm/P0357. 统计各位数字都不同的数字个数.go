/**
 * Author: Deean
 * Date: 2022/12/10 22:21
 * FileName: algorithm/P0357. 统计各位数字都不同的数字个数.go
 * Description:
 */

package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	cnt := 9
	for i := 9; i > 10-n; i-- {
		cnt *= i
	}
	return cnt + countNumbersWithUniqueDigits(n-1)
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
}
