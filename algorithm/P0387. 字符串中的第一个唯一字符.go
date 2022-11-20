/**
 * Author: Deean
 * Date: 2022/11/20 21:51
 * FileName: algorithm/P0387. 字符串中的第一个唯一字符.go
 * Description:
 */

package main

import "fmt"

func firstUniqChar(s string) int {
	hash := map[rune]int{}
	for _, c := range s {
		hash[c-'a']++
	}
	for i, c := range s {
		if hash[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
