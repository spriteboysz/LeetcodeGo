/**
 * Author: Deean
 * Date: 2022/11/20 22:49
 * FileName: algorithm/P0231. 2 的幂.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func isPowerOfTwo(n int) bool {
	return bits.OnesCount64(uint64(n)) == 1
}

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}
