/**
 * Author: Deean
 * Date: 2022-10-19 23:55
 * FileName: algorithm/P0190. 颠倒二进制位.go
 * Description:
 */

package main

import "fmt"

func reverseBits(num uint32) uint32 {
	rev := uint32(0)
	for i := 0; i < 32 && num > 0; i++ {
		rev |= num & 1 << (31 - i)
		num >>= 1
	}
	return rev
}

func main() {
	fmt.Println(reverseBits(43261596))
}
