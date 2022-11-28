/**
 * Author: Deean
 * Date: 2022/11/27 23:10
 * FileName: algorithm/P0405. 数字转换为十六进制数.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func toHex(num int) string {
	return strconv.FormatUint(uint64(uint32(num)), 16)
}
func main() {
	fmt.Println(toHex(26))
	fmt.Println(toHex(-1))
}
