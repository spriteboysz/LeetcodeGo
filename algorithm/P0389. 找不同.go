/**
 * Author: Deean
 * Date: 2022-10-21 22:11
 * FileName: algorithm/P0389. 找不同.go
 * Description:
 */

package main

import "fmt"

func findTheDifference(s string, t string) byte {
	alphabet := [26]int{}
	for _, c := range s {
		alphabet[c-'a']--
	}
	for _, c := range t {
		alphabet[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] > 0 {
			return byte(i + 'a')
		}
	}
	return byte('a')
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
