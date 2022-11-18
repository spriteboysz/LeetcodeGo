/**
 * Author: Deean
 * Date: 2022/11/17 23:23
 * FileName: algorithm/P0521. 最长特殊序列 Ⅰ.go
 * Description:
 */

package main

import "fmt"

func findLUSlength(a string, b string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if a == b {
		return -1
	} else {
		return max(len(a), len(b))
	}
}

func main() {
	fmt.Println(findLUSlength("aaa", "bbb"))
}
