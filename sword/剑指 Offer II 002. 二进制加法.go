/**
 * Author: Deean
 * Date: 2022-10-22 15:28
 * FileName: sword/剑指 Offer II 002. 二进制加法.go
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
	fmt.Println(addBinary("1010", "1011"))
}
