/**
 * Author: Deean
 * Date: 2022/11/14 23:07
 * FileName: algorithm/P1941. 检查是否所有字符出现次数相同.go
 * Description:
 */

package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	hash := make([]int, 26)
	for _, c := range s {
		hash[c-'a'] += 1
	}
	target := 0
	for i := 0; i < 26; i++ {
		if hash[i] == 0 {
			continue
		}
		if target == 0 {
			target = hash[i]
		} else if hash[i] != target {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
}
