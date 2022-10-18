/**
 * Author: Deean
 * Date: 2022-10-18 22:57
 * FileName: sword/剑指 Offer 15. 二进制中1的个数.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

func main() {
	fmt.Println(hammingWeight(11))
}
