/**
 * Author: Deean
 * Date: 2022/11/12 23:22
 * FileName: algorithm/P0506. 相对名次.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	ssp := []string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	hash, relative := make(map[int]int), make([]string, len(score))
	for i := 0; i < len(score); i++ {
		hash[score[i]] = i
	}
	sort.Slice(score, func(a, b int) bool { return score[b] < score[a] })
	for i := 0; i < len(score); i++ {
		if i <= 2 {
			relative[hash[score[i]]] = ssp[i]
		} else {
			relative[hash[score[i]]] = strconv.Itoa(i + 1)
		}
	}
	return relative
}

func main() {
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))
}
