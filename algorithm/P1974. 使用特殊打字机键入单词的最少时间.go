/**
 * Author: Deean
 * Date: 2022-11-09 22:09
 * FileName: algorithm/P1974. 使用特殊打字机键入单词的最少时间.go
 * Description:
 */

package main

import "fmt"

func minTimeToType(word string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	times, cur := 0, 'a'
	for _, c := range word {
		time := abs(int(c - cur))
		times += min(time, 26-time) + 1
		cur = c
	}
	return times
}

func main() {
	fmt.Println(minTimeToType("zjpc"))
	fmt.Println(minTimeToType("bza"))
}
