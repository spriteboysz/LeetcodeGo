/**
 * Author: Deean
 * Date: 2022-10-16 22:19
 * FileName: algorithm/P2000. 反转单词前缀.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	index := strings.Index(word, string(ch))
	if index < 0 {
		return word
	}
	var ss []byte
	for i := index; i >= 0; i-- {
		ss = append(ss, word[i])
	}
	for i := index + 1; i < len(word); i++ {
		ss = append(ss, word[i])
	}
	return string(ss)
}

func main() {
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("xyxzxe", 'a'))
	fmt.Println(reversePrefix("ayxzxe", 'a'))
	fmt.Println(reversePrefix("ayxzxe", 'e'))
}
