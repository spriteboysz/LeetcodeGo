/**
 * Author: Deean
 * Date: 2022/11/15 23:05
 * FileName: interview/面试题 01.02. 判定是否互为字符重排.go
 * Description:
 */

package main

import "fmt"

func CheckPermutation(s1 string, s2 string) bool {
	alphabet1, alphabet2 := make([]int, 26), make([]int, 26)
	for _, c := range s1 {
		alphabet1[c-'a'] += 1
	}
	for _, c := range s2 {
		alphabet2[c-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if alphabet1[i] != alphabet2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}
