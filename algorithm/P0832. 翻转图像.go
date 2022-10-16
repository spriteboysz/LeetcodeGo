/**
 * Author: Deean
 * Date: 2022-10-16 17:53
 * FileName: algorithm/P0832. 翻转图像.go
 * Description:
 */

package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	image2 := make([][]int, n)
	for i := range image2 {
		image2[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			image2[i][j] = 1 - image[i][n-1-j]
		}
	}
	return image2
}

func main() {
	image := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(flipAndInvertImage(image))
}
