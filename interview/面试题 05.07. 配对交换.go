/**
 * Author: Deean
 * Date: 2022-10-20 23:08
 * FileName: interview/面试题 05.07. 配对交换.go
 * Description:
 */

package main

import "fmt"

func exchangeBits(num int) int {
	even := num & 0xaaaaaaaa
	odd := num & 0x55555555
	return even>>1 | odd<<1
}

func main() {
	fmt.Println(exchangeBits(2))
	fmt.Println(exchangeBits(3))
}
