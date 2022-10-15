/**
 * Author: Deean
 * Date: 2022-10-15 22:11
 * FileName: algorithm/P1266. 访问所有点的最小时间.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	for i := 1; i < len(points); i++ {
		sx, sy := points[i-1][0], points[i-1][1]
		dx, dy := points[i][0], points[i][1]
		time += int(math.Max(math.Abs(float64(dx-sx)), math.Abs(float64(dy-sy))))
	}
	return time
}

func main() {
	points := [][]int{{3, 2}, {-2, 2}}
	fmt.Println(minTimeToVisitAllPoints(points))
}
