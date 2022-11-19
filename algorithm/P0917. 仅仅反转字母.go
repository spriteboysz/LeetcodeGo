/**
 * Author: Deean
 * Date: 2022/11/19 16:48
 * FileName: algorithm/P0917. 仅仅反转字母.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func reverseOnlyLetters(s string) string {
	letters := []rune{}
	for _, c := range s {
		if unicode.IsLetter(c) {
			letters = append(letters, c)
		}
	}
	reversed := []string{}
	i := len(letters) - 1
	for _, c := range s {
		if unicode.IsLetter(c) {
			reversed = append(reversed, string(letters[i]))
			i--
		} else {
			reversed = append(reversed, string(c))
		}
	}
	return strings.Join(reversed, "")
}

func main() {
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
