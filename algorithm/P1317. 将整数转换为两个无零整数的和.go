/**
 * Author: Deean
 * Date: 2022/11/19 20:39
 * FileName: algorithm/P1317. 将整数转换为两个无零整数的和.go
 * Description:
 */

package main

import "fmt"

func getNoZeroIntegers(n int) []int {
	check := func(num int) bool {
		for num > 0 {
			if num%10 == 0 {
				return false
			}
			num /= 10
		}
		return true
	}
	for i := 1; i < n/2+1; i++ {
		if check(i) && check(n-i) {
			return []int{i, n - i}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(getNoZeroIntegers(1000))
}
