/**
 * Author: Deean
 * Date: 2022-10-16 20:12
 * FileName: algorithm/P1370. 上升下降字符串.go
 * Description:
 */

package main

import (
	"fmt"
)

func sortString(s string) string {
	alphabet := [26]int{}
	for _, c := range s {
		alphabet[c-'a']++
	}
	var ss []byte
	for len(ss) < len(s) {
		for i := 0; i < 26; i++ {
			if alphabet[i] > 0 {
				ss = append(ss, byte(i+'a'))
				alphabet[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if alphabet[i] > 0 {
				ss = append(ss, byte(i+'a'))
				alphabet[i]--
			}
		}
	}
	return string(ss)
}

func main() {
	fmt.Println(sortString("leetcode"))
}
