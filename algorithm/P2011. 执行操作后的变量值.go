/**
 * Author: Deean
 * Date: 2022-10-12 23:44
 * FileName: algorithm/P2011. 执行操作后的变量值.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, operation := range operations {
		if strings.Contains(operation, "+") {
			x += 1
		} else {
			x--
		}
	}
	return x
}

func main() {
	operations := []string{"++X", "++X", "X++"}
	fmt.Println(finalValueAfterOperations(operations))
}
