/**
 * Author: Deean
 * Date: 2022-10-21 21:58
 * FileName: algorithm/P1160. 拼写单词.go
 * Description:
 */

package main

import "fmt"

func countCharacters(words []string, chars string) int {
	alphabet := [26]int{}
	for _, c := range chars {
		alphabet[c-'a']++
	}
	total := 0
	for _, word := range words {
		cur := alphabet
		for _, c := range word {
			cur[c-'a']--
		}
		flag := true
		for i := 0; i < 26; i++ {
			if cur[i] < 0 {
				flag = false
				break
			}
		}
		if flag {
			total += len(word)
		}
	}
	return total
}

func main() {
	words := []string{"hello", "world", "leetcode"}
	fmt.Println(countCharacters(words, "welldonehoneyr"))
}
