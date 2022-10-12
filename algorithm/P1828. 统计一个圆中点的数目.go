/**
 * Author: Deean
 * Date: 2022-10-12 23:10
 * FileName: algorithm/P1828. 统计一个圆中点的数目.go
 * Description:
 */

package main

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
	var count []int
	for _, query := range queries {
		a, b, r := query[0], query[1], query[2]
		cnt := 0
		for _, point := range points {
			x, y := point[0], point[1]
			if (x-a)*(x-a)+(y-b)*(y-b) <= r*r {
				cnt++
			}
		}
		count = append(count, cnt)
	}
	return count
}

func main() {
	points := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}
	queries := [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}
	fmt.Println(countPoints(points, queries))
}
