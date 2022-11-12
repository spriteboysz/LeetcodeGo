/**
 * Author: Deean
 * Date: 2022/11/12 22:56
 * FileName: algorithm/P2309. 兼具大小写的最好英文字母.go
 * Description:
 */

package main

import "fmt"

func greatestLetter(s string) string {
	hash := map[rune]bool{}
	for _, c := range s {
		hash[c] = true
	}
	for i := 'Z'; i >= 'A'; i-- {
		if hash[i] && hash[i-'A'+'a'] {
			return string(i)
		}
	}
	return ""
}

func main() {
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}
