/**
 * Author: Deean
 * Date: 2022-10-19 22:37
 * FileName: algorithm/P1047. 删除字符串中的所有相邻重复项.go
 * Description:
 */

package main

import "fmt"

func removeDuplicates(s string) string {
	var stack []byte
	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == byte(c) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
