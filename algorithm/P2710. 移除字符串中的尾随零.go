/**
 * Author: Deean
 * Date: 2023-06-24 17:11
 * FileName: algorithm/P2710. 移除字符串中的尾随零.go
 * Description:
 */

package main

import "fmt"

func removeTrailingZeros(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] != '0' {
			return num[:i+1]
		}
	}
	return num
}

func main() {
	fmt.Println(removeTrailingZeros("51230100"))
}
