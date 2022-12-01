/**
 * Author: Deean
 * Date: 2022/11/30 23:45
 * FileName: algorithm/P2347. 最好的扑克手牌.go
 * Description:
 */

package main

import (
	"bytes"
	"fmt"
)

func bestHand(ranks []int, suits []byte) string {
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	for _, r := range ranks {
		cnt[r]++
		if cnt[r] == 3 {
			return "Three of a Kind"
		}
		if cnt[r] == 2 {
			pair = true
		}
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}

func main() {
	fmt.Println(bestHand([]int{4, 4, 2, 4, 4}, []byte{'d', 'a', 'a', 'b', 'c'}))
}
