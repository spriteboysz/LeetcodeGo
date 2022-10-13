/**
 * Author: Deean
 * Date: 2022-10-13 22:46
 * FileName: algorithm/P2114. 句子中的最多单词数.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	maximum := 0
	for _, sentence := range sentences {
		cnt := len(strings.Split(sentence, " "))
		if maximum < cnt {
			maximum = cnt
		}
	}
	return maximum
}

func main() {
	sentences := []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(sentences))
}
