/**
 * Author: Deean
 * Date: 2022/11/12 20:23
 * FileName: algorithm/P0812. 最大三角形面积.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	calcArea := func(x1, y1, x2, y2, x3, y3 int) float64 {
		return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
	}
	maximum := 0.0
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				area := calcArea(p[0], p[1], q[0], q[1], r[0], r[1])
				if area > maximum {
					maximum = area
				}
			}
		}
	}
	return maximum
}

func main() {
	fmt.Println(largestTriangleArea([][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}))
}
