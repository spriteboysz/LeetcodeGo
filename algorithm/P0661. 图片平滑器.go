/**
 * Author: Deean
 * Date: 2022/11/18 23:41
 * FileName: algorithm/P0661. 图片平滑器.go
 * Description:
 */

package main

import "fmt"

func imageSmoother(img [][]int) [][]int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	m, n := len(img), len(img[0])
	smooth := make([][]int, m)
	for i := range smooth {
		smooth[i] = make([]int, n)
		for j := range smooth[i] {
			sum, num := 0, 0
			for _, row := range img[max(i-1, 0):min(j+2, n)] {
				for _, v := range row[max(j-1, 0):min(j+2, n)] {
					sum += v
					num++
				}
			}
			smooth[i][j] = sum / num
		}
	}
	return smooth
}

func main() {
	fmt.Println(imageSmoother([][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}))
}
