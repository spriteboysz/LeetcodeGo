/**
 * Author: Deean
 * Date: 2023-06-24 17:07
 * FileName: algorithm/P2716. 最小化字符串长度.go
 * Description:
 */

package main

import "fmt"

func minimizedStringLength(s string) int {
	alphabet := make([]int, 26)
	for _, c := range s {
		alphabet[c-'a'] += 1
	}
	cnt := 0
	for _, num := range alphabet {
		if num > 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(minimizedStringLength("aaabc"))
}
