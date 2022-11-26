/**
 * Author: Deean
 * Date: 2022/11/26 21:54
 * FileName: algorithm/P0520. 检测大写字母.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {
	return strings.ToUpper(word) == word || strings.ToLower(word)[1:] == word[1:]
}

func main() {
	fmt.Println(detectCapitalUse("FlaG"))
}
