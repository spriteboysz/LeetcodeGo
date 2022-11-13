/**
 * Author: Deean
 * Date: 2022/11/13 17:13
 * FileName: algorithm/P1779. 找到最近的有相同 X 或 Y 坐标的点.go
 * Description:
 */

package main

import "fmt"

func nearestValidPoint(x int, y int, points [][]int) int {
	abs := func(num int) int {
		if num >= 0 {
			return num
		}
		return -num
	}

	minIndex, minNum := -1, 10001
	for i, point := range points {
		m, n := point[0], point[1]
		if m != x && y != n {
			continue
		}
		distance := abs(x-m) + abs(y-n)
		if distance < minNum {
			minNum = distance
			minIndex = i
		}
	}
	return minIndex
}

func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
}
