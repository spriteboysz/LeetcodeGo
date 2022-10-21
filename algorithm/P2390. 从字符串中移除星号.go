/**
 * Author: Deean
 * Date: 2022-10-21 23:06
 * FileName: algorithm/P2390. 从字符串中移除星号.go
 * Description:
 */

package main

import "fmt"

func removeStars(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeStars("erase*****"))
}
