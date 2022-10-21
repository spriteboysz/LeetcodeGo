/**
 * Author: Deean
 * Date: 2022-10-21 22:55
 * FileName: algorithm/P1046. 最后一块石头的重量.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		a, b := stones[n-1], stones[n-2]
		stones = stones[:n-2]
		if a != b {
			stones = append(stones, a-b)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	} else {
		return 0
	}
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
