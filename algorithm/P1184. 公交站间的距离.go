/**
 * Author: Deean
 * Date: 2022/11/19 14:30
 * FileName: algorithm/P1184. 公交站间的距离.go
 * Description:
 */

package main

import "fmt"

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if start > destination {
		start, destination = destination, start
	}
	total, cur := 0, 0
	for i, num := range distance {
		total += num
		if i >= start && i < destination {
			cur += num
		}
	}
	return min(cur, total-cur)
}

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
}
