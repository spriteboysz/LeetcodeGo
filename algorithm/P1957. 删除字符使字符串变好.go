/**
 * Author: Deean
 * Date: 2022/12/4 21:55
 * FileName: algorithm/P1957. 删除字符使字符串变好.go
 * Description:
 */

package main

import "fmt"

func makeFancyString(s string) string {
	stack := []byte{}
	cnt := 0
	for i := range s {
		if len(stack) != 0 && s[i] == stack[len(stack)-1] {
			if cnt == 2 {
				continue
			}
			cnt++
		} else {
			cnt = 1
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}

func main() {
	fmt.Println(makeFancyString("leeetcode"))
}
