/**
 * Author: Deean
 * Date: 2022-10-20 23:42
 * FileName: algorithm/P0476. 数字的补数.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func findComplement(num int) int {
	fmt.Println(bits.Len32(uint32(num)))
	return num ^ (1<<bits.Len32(uint32(num)) - 1)
}

func main() {
	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))
}
