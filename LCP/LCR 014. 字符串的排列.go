/**
 * Author: Deean
 * Date: 2023-10-14 21:53
 * FileName: LCP/LCR 014. 字符串的排列.go
 * Description:
 */

package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, c := range s1 {
		cnt[c-'a']--
	}
	left := 0
	for right := range s2 {
		x := s2[right] - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
