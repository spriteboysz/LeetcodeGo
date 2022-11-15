/**
 * Author: Deean
 * Date: 2022/11/15 22:45
 * FileName: algorithm/P1710. 卡车上的最大单元数.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	cnt := 0
	for _, box := range boxTypes {
		if box[0] >= truckSize {
			cnt += truckSize * box[1]
			break
		}
		truckSize -= box[0]
		cnt += box[0] * box[1]
	}
	return cnt
}

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}
