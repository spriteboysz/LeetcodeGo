/**
 * Author: Deean
 * Date: 2022/12/1 22:56
 * FileName: algorithm/P2446. 判断两个事件是否存在冲突.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func haveConflict(event1 []string, event2 []string) bool {
	process := func(time string) int {
		times := strings.Split(time, ":")
		hh, _ := strconv.Atoi(times[0])
		mm, _ := strconv.Atoi(times[1])
		return hh*60 + mm
	}
	s1, e1 := process(event1[0]), process(event1[1])
	s2, e2 := process(event2[0]), process(event2[1])
	for i := 0; i < 24*60; i++ {
		if i >= s1 && i <= e1 && i >= s2 && i <= e2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(haveConflict([]string{"01:00", "02:00"}, []string{"01:20", "03:00"}))
}
