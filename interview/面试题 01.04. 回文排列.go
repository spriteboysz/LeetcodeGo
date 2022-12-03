/**
 * Author: Deean
 * Date: 2022/12/1 23:24
 * FileName: interview/面试题 01.04. 回文排列.go
 * Description:
 */

package main

import "fmt"

func canPermutePalindrome(s string) bool {
	hash := map[rune]int{}
	for _, c := range s {
		hash[c-'a']++
	}
	cnt := 0
	for _, v := range hash {
		if v%2 == 1 {
			cnt++
		}
	}
	return cnt <= 1
}

func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
}
