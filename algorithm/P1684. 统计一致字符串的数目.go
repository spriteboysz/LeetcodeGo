/**
 * Author: Deean
 * Date: 2022-10-15 23:23
 * FileName: algorithm/P1684. 统计一致字符串的数目.go
 * Description:
 */

package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	valid := make([]bool, 26)
	for _, c := range allowed {
		valid[c-'a'] = true
	}
	cnt := 0
	for _, word := range words {
		flag := true
		for _, c := range word {
			if !valid[c-'a'] {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	return cnt
}

func main() {
	words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings("cad", words))
}
