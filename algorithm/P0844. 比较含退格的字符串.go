/**
 * Author: Deean
 * Date: 2022/11/22 22:58
 * FileName: algorithm/P0844. 比较含退格的字符串.go
 * Description:
 */

package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	stack1, stack2 := []byte{}, []byte{}
	for _, c := range s {
		if c != '#' {
			stack1 = append(stack1, byte(c))
		} else if len(stack1) > 0 {
			stack1 = stack1[:len(stack1)-1]
		}
	}
	for _, c := range t {
		if c != '#' {
			stack2 = append(stack2, byte(c))
		} else if len(stack2) > 0 {
			stack2 = stack2[:len(stack2)-1]
		}
	}
	return string(stack1) == string(stack2)
}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ac#c"))
}
