/**
 * Author: Deean
 * Date: 2022-10-15 21:15
 * FileName: algorithm/P1791. 找出星型图的中心节点.go
 * Description:
 */

package main

import "fmt"

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else {
		return edges[0][1]
	}
}

func main() {
	edges := [][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}
	fmt.Println(findCenter(edges))
}
