/**
 * Author: Deean
 * Date: 2022-10-14 22:00
 * FileName: LCP/LCP 06. 拿硬币.go
 * Description:
 */

package main

import "fmt"

func minCount(coins []int) int {
	cnt := 0
	for _, coin := range coins {
		cnt += (coin + 1) / 2
	}
	return cnt
}

func main() {
	coins := []int{2, 3, 10}
	fmt.Println(minCount(coins))
}
