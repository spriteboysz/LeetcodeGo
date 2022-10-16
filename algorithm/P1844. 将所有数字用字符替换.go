/**
 * Author: Deean
 * Date: 2022-10-16 20:39
 * FileName: algorithm/P1844. 将所有数字用字符替换.go
 * Description:
 */

package main

import "fmt"

func replaceDigits(s string) string {
	var ss []byte
	for i, c := range s {
		if i%2 == 0 {
			ss = append(ss, byte(c))
		} else {
			ss = append(ss, s[i-1]+s[i]-'0')
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(replaceDigits("a1b2c3d4e"))
}
