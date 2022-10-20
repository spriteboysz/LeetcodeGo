/**
 * Author: Deean
 * Date: 2022-10-20 23:09
 * FileName: algorithm/P2068. 检查两个字符串是否几乎相等.go
 * Description:
 */

package main

import "fmt"

func checkAlmostEquivalent(word1 string, word2 string) bool {
	alphabet := [26]int{}
	for _, c := range word1 {
		alphabet[c-'a']++
	}
	for _, c := range word2 {
		alphabet[c-'a']--
	}

	abs := func(num int) int {
		if num < 0 {
			num = -num
		}
		return num
	}

	for _, cnt := range alphabet {
		if abs(cnt) > 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkAlmostEquivalent("abcdeef", "abaaacc"))
	fmt.Println(checkAlmostEquivalent("cccddabba", "babababab"))
}
