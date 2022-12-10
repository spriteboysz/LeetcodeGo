/**
 * Author: Deean
 * Date: 2022/12/9 23:25
 * FileName: algorithm/P0648. 单词替换.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	// fmt.Println(dictionary)
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for _, dict := range dictionary {
			if strings.HasPrefix(word, dict) {
				words[i] = dict
				break
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))
}
