/**
 * Author: Deean
 * Date: 2022-10-13 00:03
 * FileName: sword/剑指 Offer 64. 求1+2+…+n.go
 * Description:
 */

package main

import "fmt"

func sumNums(n int) int {
	if n == 1 {
		return 1
	}
	return sumNums(n-1) + n
}

func main() {
	fmt.Println(sumNums(100))
}
