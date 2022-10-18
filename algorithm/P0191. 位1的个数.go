/**
 * Author: Deean
 * Date: 2022-10-17 23:49
 * FileName: algorithm/P0191. 位1的个数.go
 * Description:
 */

package main

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num = num & (num - 1)
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(hammingWeight(3))
	fmt.Println(hammingWeight(4))
}
