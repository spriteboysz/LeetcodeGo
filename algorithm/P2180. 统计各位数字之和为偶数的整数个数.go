/**
 * Author: Deean
 * Date: 2022-10-21 23:30
 * FileName: algorithm/P2180. 统计各位数字之和为偶数的整数个数.go
 * Description:
 */

package main

import "fmt"

func countEven(num int) int {
	process := func(num int) (sum int) {
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		return
	}
	cnt := 0
	for i := 2; i <= num; i++ {
		if process(i)%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countEven(30))
}
