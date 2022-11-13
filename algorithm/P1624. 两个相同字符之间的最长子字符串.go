/**
 * Author: Deean
 * Date: 2022/11/13 17:58
 * FileName: algorithm/P1624. 两个相同字符之间的最长子字符串.go
 * Description:
 */

package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	alphabet := make([]int, 26)
	for i := range alphabet {
		alphabet[i] = -1
	}
	maximum := -1
	for i, c := range s {
		c -= 'a'
		if alphabet[c] == -1 {
			alphabet[c] = i
		} else {
			maximum = max(maximum, i-alphabet[c]-1)
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("cbzxy"))
	fmt.Println(maxLengthBetweenEqualCharacters("cabbac"))
}
