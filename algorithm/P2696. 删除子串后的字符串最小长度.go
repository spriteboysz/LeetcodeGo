/**
 * Author: Deean
 * Date: 2023-06-24 20:37
 * FileName: algorithm/P2696. 删除子串后的字符串最小长度.go
 * Description:
 */

package main

import "fmt"

func minLength(s string) int {
	stack := []byte{}
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, byte(c))
		} else if stack[len(stack)-1] == 'A' && c == 'B' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == 'C' && c == 'D' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(c))
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(minLength("ABFCACDB"))
}
