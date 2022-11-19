/**
 * Author: Deean
 * Date: 2022/11/19 16:37
 * FileName: sword/P0242. 有效的字母异位词.go
 * Description:
 */

package main

import "fmt"

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alphabet := [26]int{}
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram2("anagram", "nagaram"))
}
