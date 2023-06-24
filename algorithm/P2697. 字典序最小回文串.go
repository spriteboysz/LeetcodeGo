/**
 * Author: Deean
 * Date: 2023-06-24 20:32
 * FileName: algorithm/P2697. 字典序最小回文串.go
 * Description:
 */

package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	min := func(a, b byte) byte {
		if a < b {
			return a
		}
		return b
	}
	ss := []byte(s)
	for i, n := 0, len(ss); i < n/2; i++ {
		c := min(ss[i], ss[n-i-1])
		ss[i] = c
		ss[n-i-1] = c
	}
	return string(ss)
}

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
}
