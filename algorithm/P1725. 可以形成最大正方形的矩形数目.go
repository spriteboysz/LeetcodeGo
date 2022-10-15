/**
 * Author: Deean
 * Date: 2022-10-15 21:50
 * FileName: algorithm/P1725. 可以形成最大正方形的矩形数目.go
 * Description:
 */

package main

import "fmt"

func countGoodRectangles(rectangles [][]int) int {
	var sides []int
	maximum := 0
	for i, rectangle := range rectangles {
		a, b := rectangle[0], rectangle[1]
		if a > b {
			sides = append(sides, b)
		} else {
			sides = append(sides, a)
		}
		if sides[i] > maximum {
			maximum = sides[i]
		}
	}
	cnt := 0
	for _, side := range sides {
		if side == maximum {
			cnt++
		}
	}
	return cnt
}

func main() {
	rectangles := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}
	fmt.Println(countGoodRectangles(rectangles))
}
