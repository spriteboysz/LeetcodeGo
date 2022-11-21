/**
 * Author: Deean
 * Date: 2022/11/21 23:02
 * FileName: algorithm/P1576. 替换所有的问号.go
 * Description:
 */

package main

import "fmt"

func modifyString(s string) string {
	chars := make([]byte, len(s))
	for i := 0; i < len(chars); i++ {
		if s[i] != '?' {
			chars[i] = s[i]
			continue
		}
		for j := byte('a'); j <= 'z'; j++ {
			if (i == 0 || j != chars[i-1]) && (i == len(chars)-1 || j != s[i+1]) {
				chars[i] = j
				break
			}
		}
	}
	return string(chars)
}

func main() {
	fmt.Println(modifyString("ubv?w"))
}
