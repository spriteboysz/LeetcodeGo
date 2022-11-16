/**
 * Author: Deean
 * Date: 2022/11/16 23:59
 * FileName: sword/剑指 Offer 50. 第一个只出现一次的字符.go
 * Description:
 */

package main

import "fmt"

func firstUniqChar(s string) byte {
	hash := map[rune]int{}
	for _, c := range s {
		hash[c]++
	}
	for _, c := range s {
		if hash[c] == 1 {
			return byte(c)
		}
	}
	return ' '
}

func main() {
	fmt.Println(firstUniqChar("abaccdeff"))
}
