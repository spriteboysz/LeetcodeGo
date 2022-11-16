/**
 * Author: Deean
 * Date: 2022/11/16 22:57
 * FileName: algorithm/P0796. 旋转字符串.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}

func main() {
	fmt.Println(rotateString("abc", "bca"))
}
