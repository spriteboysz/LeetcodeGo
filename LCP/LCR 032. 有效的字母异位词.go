/**
 * Author: Deean
 * Date: 2023-10-15 10:07
 * FileName: LCP/LCR 032. 有效的字母异位词.go
 * Description:
 */

package main

import "fmt"

func isAnagram(s string, t string) bool {
	m, n := len(s), len(t)
	if m != n || s == t {
		return false
	}
	alphabet := [26]int{}
	for i := range s {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}
	for _, num := range alphabet {
		if num != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
