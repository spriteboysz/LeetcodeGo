/**
 * Author: Deean
 * Date: 2023-10-14 20:23
 * FileName: LCP/LCR 002. 二进制求和.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	sum, carry := "", 0
	for pos1, pos2 := len(a)-1, len(b)-1; pos1 >= 0 || pos2 >= 0 || carry > 0; pos1, pos2 = pos1-1, pos2-1 {
		if pos1 >= 0 {
			carry += int(a[pos1] - '0')
		}
		if pos2 >= 0 {
			carry += int(b[pos2] - '0')
		}
		sum = strconv.Itoa(carry%2) + sum
		carry >>= 1
	}
	return sum
}

func main() {
	fmt.Println(addBinary("11", "1111"))
}
