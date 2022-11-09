/**
 * Author: Deean
 * Date: 2022-11-09 22:45
 * FileName: algorithm/P1742. 盒子中小球的最大数量.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countBalls(lowLimit int, highLimit int) int {
	count := [50]int{}
	for i := lowLimit; i <= highLimit; i++ {
		cur := 0
		for _, c := range strings.Split(strconv.Itoa(i), "") {
			num, _ := strconv.Atoi(c)
			cur += num
		}
		count[cur] += 1
	}
	maximum := 0
	for _, c := range count {
		if c > maximum {
			maximum = c
		}
	}
	return maximum
}

func main() {
	fmt.Println(countBalls(5, 15))
	fmt.Println(countBalls(19, 28))
}
