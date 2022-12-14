/**
 * Author: Deean
 * Date: 2022/12/13 23:28
 * FileName: algorithm/P1897. 重新分配字符使所有字符串都相等.go
 * Description:
 */

package main

import "fmt"

func makeEqual(words []string) bool {
	alphabet := [26]int{}
	for _, word := range words {
		for _, c := range word {
			alphabet[c-'a']++
		}
	}
	for _, v := range alphabet {
		if v%len(words) != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(makeEqual([]string{"abc", "aabc", "bc"}))
}
