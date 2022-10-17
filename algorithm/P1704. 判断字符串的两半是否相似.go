/**
 * Author: Deean
 * Date: 2022-10-17 22:49
 * FileName: algorithm/P1704. 判断字符串的两半是否相似.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	cnt := 0
	for i, c := range s {
		if strings.Contains(vowels, string(c)) {
			if i < len(s)/2 {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cnt == 0
}

func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}
