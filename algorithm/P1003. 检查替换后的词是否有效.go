/**
 * Author: Deean
 * Date: 2022/12/11 17:33
 * FileName: algorithm/P1003. 检查替换后的词是否有效.go
 * Description:
 */

package main

import "fmt"

func isValid2(s string) bool {
	stack := []byte{}
	for _, c := range s {
		if c != 'c' {
			stack = append(stack, byte(c))
		} else {
			if len(stack) >= 2 && stack[len(stack)-1] == 'b' && stack[len(stack)-2] == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid2("abcabcababcc"))
}
