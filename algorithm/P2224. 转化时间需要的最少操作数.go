/**
 * Author: Deean
 * Date: 2022-10-21 21:22
 * FileName: algorithm/P2224. 转化时间需要的最少操作数.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	process := func(time string) int {
		times := strings.Split(time, ":")
		hh, _ := strconv.Atoi(times[0])
		mm, _ := strconv.Atoi(times[1])
		return hh*60 + mm
	}
	diff := process(correct) - process(current)
	cnt := 0
	for _, item := range []int{60, 15, 5, 1} {
		cnt += diff / item
		diff %= item
	}
	return cnt
}

func main() {
	fmt.Println(convertTime("02:30", "04:35"))
}
