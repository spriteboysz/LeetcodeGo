/**
 * Author: Deean
 * Date: 2022-10-14 23:56
 * FileName: algorithm/P1678. 设计 Goal 解析器.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	return strings.ReplaceAll(command, "()", "o")
}

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}
