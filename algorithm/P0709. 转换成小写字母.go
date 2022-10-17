/**
 * Author: Deean
 * Date: 2022-10-17 23:03
 * FileName: algorithm/P0709. 转换成小写字母.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	return strings.ToLower(s)
}

func main() {
	fmt.Println(toLowerCase("LOVELY"))
}
