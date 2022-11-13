/**
 * Author: Deean
 * Date: 2022/11/12 23:48
 * FileName: algorithm/P0693. 交替位二进制数.go
 * Description:
 */

package main

import "fmt"

func hasAlternatingBits(n int) bool {
	for pre := 2; n != 0; n /= 2 {
		cur := n % 2
		if cur == pre {
			return false
		}
		pre = cur
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(7))
}
