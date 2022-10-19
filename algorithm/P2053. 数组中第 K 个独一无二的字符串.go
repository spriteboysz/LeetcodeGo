/**
 * Author: Deean
 * Date: 2022-10-19 23:25
 * FileName: algorithm/P2053. 数组中第 K 个独一无二的字符串.go
 * Description:
 */

package main

import "fmt"

func kthDistinct(arr []string, k int) string {
	hash := map[string]int{}
	for _, word := range arr {
		hash[word]++
	}
	for _, word := range arr {
		if hash[word] == 1 {
			k--
		}
		if k == 0 {
			return word
		}
	}
	return ""
}

func main() {
	arr := []string{"aaa", "aa", "a"}
	fmt.Println(kthDistinct(arr, 1))
	fmt.Println(kthDistinct(arr, 4))
}
