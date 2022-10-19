/**
 * Author: Deean
 * Date: 2022-10-20 00:04
 * FileName: interview/面试题 01.01. 判定字符是否唯一.go
 * Description:
 */

package main

import "fmt"

func isUnique(astr string) bool {
	hash := map[byte]bool{}
	for _, c := range astr {
		if hash[byte(c)] == true {
			return false
		} else {
			hash[byte(c)] = true
		}
	}
	return true
}

func main() {
	fmt.Println(isUnique("leetcode"))
}
