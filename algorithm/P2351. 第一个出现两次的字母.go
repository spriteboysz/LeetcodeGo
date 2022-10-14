/**
 * Author: Deean
 * Date: 2022-10-14 21:31
 * FileName: algorithm/P2351. 第一个出现两次的字母.go
 * Description:
 */

package main

import "fmt"

func repeatedCharacter(s string) byte {
	seen := make(map[rune]bool, 0)
	for _, c := range s {
		if seen[c] {
			return byte(c)
		} else {
			seen[c] = true
		}
	}
	return byte('a')
}

func main() {
	fmt.Println(repeatedCharacter("abcdd"))
}
