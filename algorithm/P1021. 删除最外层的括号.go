/**
 * Author: Deean
 * Date: 2022-10-16 12:14
 * FileName: algorithm/P1021. 删除最外层的括号.go
 * Description:
 */

package main

import "fmt"

func removeOuterParentheses(s string) string {
	var ss []rune
	var cur []rune
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		cur = append(cur, c)
		if balance == 0 {
			ss = append(ss, cur[1:len(cur)-1]...)
			cur = []rune{}
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
