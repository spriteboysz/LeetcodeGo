/**
 * Author: Deean
 * Date: 2022-10-12 23:51
 * FileName: sword/剑指 Offer 58 - II. 左旋转字符串.go
 * Description:
 */

package main

import (
	"fmt"
)

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func main() {
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}
