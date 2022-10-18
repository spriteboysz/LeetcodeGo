/**
 * Author: Deean
 * Date: 2022-10-18 23:33
 * FileName: algorithm/P2119. 反转两次的数字.go
 * Description:
 */

package main

import "fmt"

func isSameAfterReversals(num int) bool {
	return num == 0 || num%10 != 0
}

func main() {
	fmt.Println(isSameAfterReversals(1800))
	fmt.Println(isSameAfterReversals(0))
}
