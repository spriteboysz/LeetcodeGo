/**
 * Author: Deean
 * Date: 2022/12/14 22:18
 * FileName: algorithm/P1128. 等价多米诺骨牌对的数量.go
 * Description:
 */

package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	nums := [100]int{}
	cnt := 0
	for _, domino := range dominoes {
		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		v := domino[0]*10 + domino[1]
		cnt += nums[v]
		nums[v]++
	}
	return cnt
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}))
}
