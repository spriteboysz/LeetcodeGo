/**
 * Author: Deean
 * Date: 2022-10-18 00:07
 * FileName: sword/剑指 Offer 05. 替换空格.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
