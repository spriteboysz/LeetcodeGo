/**
 * Author: Deean
 * Date: 2022/11/16 22:06
 * FileName: algorithm/P1903. 字符串中的最大奇数.go
 * Description:
 */

package main

import "fmt"

func largestOddNumber(num string) string {
	for num != "" && num[len(num)-1]&1 == 0 {
		num = num[:len(num)-1]
	}
	return num
}

func main() {
	fmt.Println(largestOddNumber("35427"))
}
