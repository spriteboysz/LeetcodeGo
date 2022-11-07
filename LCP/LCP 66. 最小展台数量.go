/**
 * Author: Deean
 * Date: 2022-11-07 23:07
 * FileName: LCP/LCP 66. 最小展台数量.go
 * Description:
 */

package main

import "fmt"

func minNumBooths(demand []string) int {
	alphabet := [26]int{}
	for _, s := range demand {
		cur := [26]int{}
		for _, c := range s {
			cur[c-'a'] += 1
		}
		for i, c := range cur {
			if c > alphabet[i] {
				alphabet[i] = c
			}
		}
	}
	cnt := 0
	for _, c := range alphabet {
		cnt += c
	}
	return cnt
}

func main() {
	fmt.Println(minNumBooths([]string{"acd", "bed", "accd"}))
}
