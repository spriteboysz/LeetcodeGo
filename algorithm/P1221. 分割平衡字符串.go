/**
 * Author: Deean
 * Date: 2022-10-13 23:59
 * FileName: algorithm/P1221. 分割平衡字符串.go
 * Description:
 */

package main

import "fmt"

func balancedStringSplit(s string) int {
	balance, cur := 0, 0
	for _, c := range s {
		if c == 'R' {
			cur++
		} else {
			cur--
		}
		if cur == 0 {
			balance++
		}
	}
	return balance
}

func main() {
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
}
