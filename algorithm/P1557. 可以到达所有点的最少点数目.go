/**
 * Author: Deean
 * Date: 2022-10-16 16:03
 * FileName: algorithm/P1557. 可以到达所有点的最少点数目.go
 * Description:
 */

package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	end := make([]int, n)
	for _, edge := range edges {
		end[edge[1]] = 1
	}
	var res []int
	for i := 0; i < n; i++ {
		if end[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	edges := [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}
	fmt.Println(findSmallestSetOfVertices(5, edges))
}
