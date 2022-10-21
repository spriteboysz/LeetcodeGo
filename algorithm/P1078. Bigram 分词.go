/**
 * Author: Deean
 * Date: 2022-10-21 23:23
 * FileName: algorithm/P1078. Bigram 分词.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	target := []string{}
	if len(words) < 3 {
		return target
	}
	for i := 2; i < len(words); i++ {
		if words[i-2] == first && words[i-1] == second {
			target = append(target, words[i])
		}
	}
	return target
}

func main() {
	fmt.Println(findOcurrences("we will we will rock you", "we", "will"))
}
