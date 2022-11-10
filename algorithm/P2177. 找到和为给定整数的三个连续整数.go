/**
 * Author: Deean
 * Date: 2022-11-10 22:57
 * FileName: algorithm/P2177. 找到和为给定整数的三个连续整数.go
 * Description:
 */

package main

import "fmt"

func sumOfThree(num int64) []int64 {
	if num%3 == 0 {
		return []int64{num/3 - 1, num / 3, num/3 + 1}
	}
	return []int64{}
}

func main() {
	fmt.Println(sumOfThree(33))
}
