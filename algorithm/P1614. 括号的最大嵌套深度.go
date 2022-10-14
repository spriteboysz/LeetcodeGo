/**
 * Author: Deean
 * Date: 2022-10-15 00:00
 * FileName: algorithm/P1614. 括号的最大嵌套深度.go
 * Description:
 */

package main

import "fmt"

func maxDepth(s string) int {
	maximum, cur := 0, 0
	for _, c := range s {
		if c == '(' {
			cur++
		} else if c == ')' {
			cur--
		}
		if cur > maximum {
			maximum = cur
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
}
