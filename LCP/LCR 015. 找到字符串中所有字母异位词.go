/**
 * Author: Deean
 * Date: 2023-10-14 22:03
 * FileName: LCP/LCR 015. 找到字符串中所有字母异位词.go
 * Description:
 */

package main

import "fmt"

func findAnagrams(s string, p string) []int {
	index, m, n := []int{}, len(s), len(p)
	if m < n {
		return index
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range p {
		cnt1[s[i]-'a']++
		cnt2[p[i]-'a']++
	}
	if cnt1 == cnt2 {
		index = append(index, 0)
	}
	for i := range s[n:] {
		cnt1[s[i]-'a']--
		cnt1[s[i+n]-'a']++
		if cnt1 == cnt2 {
			index = append(index, i+1)
		}
	}
	return index
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("aa", "aaa"))
}
