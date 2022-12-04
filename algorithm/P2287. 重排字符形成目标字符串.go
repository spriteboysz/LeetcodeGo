/**
 * Author: Deean
 * Date: 2022/12/4 21:48
 * FileName: algorithm/P2287. 重排字符形成目标字符串.go
 * Description:
 */

package main

import "fmt"

func rearrangeCharacters(s string, target string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for _, c := range s {
		cnt1[c-'a']++
	}
	for _, c := range target {
		cnt2[c-'a']++
	}
	cnt := len(s)
	for i, c := range cnt1 {
		if cnt2[i] > 0 {
			cnt = min(cnt, c/cnt2[i])
		}
	}
	return cnt
}

func main() {
	fmt.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}
