/**
 * Author: Deean
 * Date: 2022/12/15 23:03
 * FileName: algorithm/P2249. 统计圆内格点数目.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func countLatticePoints(circles [][]int) int {
	sort.Slice(circles, func(i, j int) bool {
		return circles[i][2] > circles[j][2]
	})

	cnt := 0
	for x := 0; x <= 200; x++ {
		for y := 0; y <= 200; y++ {
			for _, circle := range circles {
				if (x-circle[0])*(x-circle[0])+(y-circle[1])*(y-circle[1]) <= circle[2]*circle[2] {
					cnt++
					break
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
}
