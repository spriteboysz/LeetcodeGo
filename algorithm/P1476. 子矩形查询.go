/**
 * Author: Deean
 * Date: 2022-11-07 22:52
 * FileName: algorithm/P1476. 子矩形查询.go
 * Description:
 */

package main

import "fmt"

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for row := row1; row <= row2; row++ {
		for col := col1; col <= col2; col++ {
			this.rectangle[row][col] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}

func main() {
	rectangle := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	subrectangleQueries := Constructor(rectangle)
	fmt.Println(subrectangleQueries.GetValue(0, 0)) // 返回 1
	subrectangleQueries.UpdateSubrectangle(0, 0, 2, 2, 100)
	fmt.Println(subrectangleQueries.GetValue(0, 0)) // 返回 100
	fmt.Println(subrectangleQueries.GetValue(2, 2)) // 返回 100
	subrectangleQueries.UpdateSubrectangle(1, 1, 2, 2, 20)
	fmt.Println(subrectangleQueries.GetValue(2, 2)) // 返回 20
}
