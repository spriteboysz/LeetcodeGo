/**
 * Author: Deean
 * Date: 2022/11/19 16:24
 * FileName: algorithm/P0492. 构造矩形.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	width := int(math.Sqrt(float64(area)))
	for i := width; i > 0; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}

func main() {
	fmt.Println(constructRectangle(122122))
}
