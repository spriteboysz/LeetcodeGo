/**
 * Author: Deean
 * Date: 2023-10-15 15:44
 * FileName: LCP/LCR 035. 最小时间差.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
	"sort"
)

func findMinDifference(timePoints []string) int {
	Min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	process := func(time string) int {
		return int(time[0]-'0')*600 + int(time[1]-'0')*60 + int(time[3]-'0')*10 + int(time[4]-'0')
	}

	sort.Strings(timePoints)
	minimum := math.MaxInt32
	t0Minutes := process(timePoints[0])
	preMinutes := t0Minutes
	for _, t := range timePoints[1:] {
		minutes := process(t)
		minimum = Min(minimum, minutes-preMinutes)
		preMinutes = minutes
	}
	minimum = Min(minimum, t0Minutes+1440-preMinutes)
	return minimum
}

func main() {
	fmt.Println(findMinDifference([]string{"00:00", "23:59", "00:00"}))
}
