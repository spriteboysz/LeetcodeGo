/**
 * Author: Deean
 * Date: 2022/11/16 23:47
 * FileName: algorithm/P2264. 字符串中最大的 3 位相同数字.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func largestGoodInteger(num string) string {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maximum := -1
	for i := 2; i < len(num); i++ {
		if num[i] == num[i-1] && num[i] == num[i-2] {
			maximum = max(maximum, int(num[i]-'0'))
		}
		if maximum == 9 {
			return "999"
		}
	}
	if maximum == -1 {
		return ""
	}
	str := strconv.Itoa(maximum)
	return str + str + str
}

func main() {
	fmt.Println(largestGoodInteger("6777133339"))
	fmt.Println(largestGoodInteger("2300019"))
}
