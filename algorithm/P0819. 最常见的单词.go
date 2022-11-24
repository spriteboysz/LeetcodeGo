/**
 * Author: Deean
 * Date: 2022/11/23 23:09
 * FileName: algorithm/P0819. 最常见的单词.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]bool{}
	for _, s := range banned {
		ban[s] = true
	}
	hash := map[string]int{}
	s := strings.ToLower(paragraph)
	r := strings.NewReplacer(";", " ", "!", " ", "?", " ", ",", " ", ".", " ", "'", " ")
	s = r.Replace(s)
	words := strings.Fields(s)
	for _, word := range words {
		if !ban[word] {
			hash[word]++
		}
	}
	maximum, common := 0, ""
	for k, v := range hash {
		if maximum < v {
			maximum = v
			common = k
		}
	}
	return common
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.",
		[]string{"hit"}))
}
