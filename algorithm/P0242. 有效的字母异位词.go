/**
 * Author: Deean
 * Date: 2022-10-21 23:15
 * FileName: algorithm/P0242. 有效的字母异位词.go
 * Description:
 */

package main

import "fmt"

func isAnagram(s string, t string) bool {
	alphabet := [26]int{}
	for _, c := range s {
		alphabet[c-'a']++
	}
	for _, c := range t {
		alphabet[c-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
