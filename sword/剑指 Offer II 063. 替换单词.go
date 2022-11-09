/**
 * Author: Deean
 * Date: 2022-11-09 23:58
 * FileName: sword/剑指 Offer II 063. 替换单词.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	hash := map[string]bool{}
	for _, d := range dictionary {
		hash[d] = true
	}
	words := strings.Split(sentence, " ")

	for i, word := range words {
		for j := 1; j < len(word); j++ {
			if hash[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"a", "b", "c"},
		"aadsfasf absbs bbab cadsfafs"))
}
