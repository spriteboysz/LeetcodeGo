/**
 * Author: Deean
 * Date: 2022-10-18 23:42
 * FileName: algorithm/P0500. 键盘行.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	hash := map[string]int{
		"qwertyuiop": 1,
		"asdfghjkl":  2,
		"zxcvbnm":    4,
	}
	var target []string
	for _, word := range words {
		cur := 0
		for _, c := range word {
			for k, v := range hash {
				if strings.Contains(k, strings.ToLower(string(c))) {
					cur |= v
				}
			}
		}
		if cur == 1 || cur == 2 || cur == 4 {
			target = append(target, word)
		}
	}
	return target
}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}
