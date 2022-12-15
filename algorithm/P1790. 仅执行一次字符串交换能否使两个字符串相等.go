/**
 * Author: Deean
 * Date: 2022/12/15 22:53
 * FileName: algorithm/P1790. 仅执行一次字符串交换能否使两个字符串相等.go
 * Description:
 */

package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	i, j := -1, -1
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			if i < 0 {
				i = idx
			} else if j < 0 {
				j = idx
			} else {
				return false
			}
		}
	}
	return i < 0 || j >= 0 && s1[i] == s2[j] && s1[j] == s2[i]
}

func main() {
	fmt.Println(areAlmostEqual("bank", "kanb"))
}
