/**
 * Author: Deean
 * Date: 2022/11/19 16:40
 * FileName: algorithm/P0495. 提莫攻击.go
 * Description:
 */

package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	times, expired := 0, 0
	for _, t := range timeSeries {
		if t > expired {
			times += duration
		} else {
			times += duration + t - expired
		}
		expired = t + duration
	}
	return times
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2))
}
