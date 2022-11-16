/**
 * Author: Deean
 * Date: 2022/11/16 22:11
 * FileName: algorithm/P1446. 连续字符.go
 * Description:
 */

package main

import "fmt"

func maxPower(s string) int {
	cnt := make([]int, len(s))
	for i := range cnt {
		cnt[i] = 1
	}
	maximum := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt[i] += cnt[i-1]
		}
		if maximum < cnt[i] {
			maximum = cnt[i]
		}
	}
	return maximum
}

func main() {
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
}
