/**
 * Author: Deean
 * Date: 2022/11/14 23:34
 * FileName: algorithm/P2129. 将标题首字母大写.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	for i, word := range words {
		word = strings.ToLower(word)
		if len(word) > 2 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(capitalizeTitle("capiTalIze tHe titLe"))
}
