/**
 * Author: Deean
 * Date: 2022-10-22 15:13
 * FileName: algorithm/P0415. 字符串相加.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	sum, carry := "", 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(num1[i] - '0')
		}
		if j >= 0 {
			carry += int(num2[j] - '0')
		}
		sum = strconv.Itoa(carry%10) + sum
		carry /= 10
	}
	return sum
}

func main() {
	fmt.Println(addStrings("456", "77"))
}
