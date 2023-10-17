/**
 * Author: Deean
 * Date: 2023-10-15 23:31
 * FileName: LCP/LCR 121. 寻找目标值 - 二维数组.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func findTargetIn2DPlants(plants [][]int, target int) bool {
	for _, plant := range plants {
		for _, num := range plant {
			if num == target {
				return true
			}
		}
	}
	return false
}

func main() {
	plants := lib.Str2Array2D("[[2,3,6,8],[4,5,8,9],[5,9,10,12]]")
	fmt.Println(findTargetIn2DPlants(plants, 8))
}
