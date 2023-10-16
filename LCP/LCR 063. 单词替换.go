/**
 * Author: Deean
 * Date: 2023-10-15 22:39
 * FileName: LCP/LCR 063. 单词替换.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	hash := map[string]bool{}
	for _, root := range dictionary {
		hash[root] = true
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
	dictionary := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(replaceWords(dictionary, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"))
}
