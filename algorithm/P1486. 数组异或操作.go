/**
 * Author: Deean
 * Date: 2022-10-13 22:58
 * FileName: algorithm/P1486. 数组异或操作.go
 * Description:
 */

package main

import "fmt"

func xorOperation(n int, start int) int {
	ret := 0
	for i := 0; i < n; i++ {
		ret ^= start + 2*i
	}
	return ret
}

func main() {
	fmt.Println(xorOperation(4, 3))  // 8
	fmt.Println(xorOperation(10, 5)) // 2
}
