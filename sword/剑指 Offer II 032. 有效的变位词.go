/**
 * Author: Deean
 * Date: 2022/11/19 16:33
 * FileName: sword/剑指 Offer II 032. 有效的变位词.go
 * Description:
 */

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	alphabet := [26]int{}
	for i, c := range s {
		alphabet[c-'a']++
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
	fmt.Println(isAnagram("anagram", "nagaram"))
}
