/**
 * Author: Deean
 * Date: 2022/12/14 22:59
 * FileName: algorithm/P1451. 重新排列句子中的单词.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	words := strings.Split(strings.ToLower(text), " ")
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	words[0] = strings.ToTitle(string(words[0][0])) + words[0][1:]
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(arrangeWords("Keep calm and code on"))
}
