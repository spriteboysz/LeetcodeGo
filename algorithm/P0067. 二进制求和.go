/**
 * Author: Deean
 * Date: 2022-10-22 15:31
 * FileName: interview/P0067. 二进制求和.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	sum, carry := "", 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(a[i] - '0')
		}
		if j >= 0 {
			carry += int(b[j] - '0')
		}
		sum = strconv.Itoa(carry%2) + sum
		carry >>= 1
	}
	return sum
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
