/**
 * Author: Deean
 * Date: 2022-11-09 23:47
 * FileName: algorithm/P0575. 分糖果.go
 * Description:
 */

package main

import "fmt"

func distributeCandies(candyType []int) int {
	hash := map[int]bool{}
	for _, t := range candyType {
		hash[t] = true
	}
	m, n := len(hash), len(candyType)/2
	if m < n {
		return m
	}
	return n
}

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(distributeCandies([]int{6, 6, 6, 6}))
}
