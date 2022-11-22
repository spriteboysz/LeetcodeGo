/**
 * Author: Deean
 * Date: 2022/11/21 23:53
 * FileName: algorithm/P0342. 4的幂.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func isPowerOfFour(n int) bool {
	return bits.OnesCount32(uint32(n)) == 1 && (bits.Len32(uint32(n))-1)%2 == 0
}

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(1))
}
