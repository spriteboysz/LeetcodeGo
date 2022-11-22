/**
 * Author: Deean
 * Date: 2022/11/22 22:55
 * FileName: algorithm/P1037. 有效的回旋镖.go
 * Description:
 */

package main

import "fmt"

func isBoomerang(points [][]int) bool {
	v1 := [2]int{points[1][0] - points[0][0], points[1][1] - points[0][1]}
	v2 := [2]int{points[2][0] - points[0][0], points[2][1] - points[0][1]}
	return v1[0]*v2[1] != v1[1]*v2[0]
}

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
}
