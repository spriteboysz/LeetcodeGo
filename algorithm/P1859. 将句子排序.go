/**
 * Author: Deean
 * Date: 2022-11-07 23:48
 * FileName: algorithm/P1859. 将句子排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	sort.Slice(words, func(i, j int) bool {
		return words[i][len(words[i])-1] < words[j][len(words[j])-1]
	})
	for i := range words {
		words[i] = words[i][:len(words[i])-1]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}
