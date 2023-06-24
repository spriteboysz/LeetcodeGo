/**
 * Author: Deean
 * Date: 2023-06-24 20:46
 * FileName: algorithm/P2682. 找出转圈游戏输家.go
 * Description:
 */

package main

import "fmt"

func circularGameLosers(n int, k int) []int {
	visited := make([]bool, n)
	for i, d := 0, k; !visited[i]; d += k {
		visited[i] = true
		i = (i + d) % n
	}

	losers := []int{}
	for i, item := range visited {
		if !item {
			losers = append(losers, i+1)
		}
	}
	return losers
}

func main() {
	fmt.Println(circularGameLosers(5, 2))
}
