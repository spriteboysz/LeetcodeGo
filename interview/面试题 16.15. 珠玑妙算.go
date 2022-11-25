/**
 * Author: Deean
 * Date: 2022/11/25 22:59
 * FileName: interview/面试题 16.15. 珠玑妙算.go
 * Description:
 */

package main

import "fmt"

func masterMind(solution string, guess string) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	hash1, hash2 := map[uint8]int{}, map[uint8]int{}
	a, b := 0, 0
	for i := 0; i < 4; i++ {
		s, g := solution[i], guess[i]
		if s == g {
			a++
		} else {
			hash1[s-'A']++
			hash2[g-'A']++
		}
	}
	for k, v := range hash1 {
		b += min(v, hash2[k])
	}
	return []int{a, b}
}

func main() {
	fmt.Println(masterMind("RGBY", "GGRR"))
}
