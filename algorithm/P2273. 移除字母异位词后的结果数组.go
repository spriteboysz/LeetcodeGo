/**
 * Author: Deean
 * Date: 2022/11/18 23:56
 * FileName: algorithm/P2273. 移除字母异位词后的结果数组.go
 * Description:
 */

package main

import "fmt"

func removeAnagrams(words []string) []string {
	ans := []string{words[0]}
	for _, word := range words[1:] {
		cnt := [26]int{}
		for _, c := range word {
			cnt[c-'a']++
		}
		for _, c := range ans[len(ans)-1] {
			cnt[c-'a']--
		}
		if cnt != [26]int{} {
			ans = append(ans, word)
		}
	}
	return ans
}

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
}
