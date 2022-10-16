/**
 * Author: Deean
 * Date: 2022-10-16 14:28
 * FileName: algorithm/P1342. 将数字变成 0 的操作次数.go
 * Description:
 */

package main

import "fmt"

func numberOfSteps(num int) int {
	cnt := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(numberOfSteps(123))
}
