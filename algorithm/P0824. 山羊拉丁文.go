/**
 * Author: Deean
 * Date: 2022/11/12 23:39
 * FileName: algorithm/P0824. 山羊拉丁文.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	vowels := "aeiouAEIOU"
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.Contains(vowels, string(word[0])) {
			words[i] += "ma"
		} else {
			words[i] = word[1:] + string(word[0]) + "ma"
		}
		for j := 0; j <= i; j++ {
			words[i] += "a"
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}
