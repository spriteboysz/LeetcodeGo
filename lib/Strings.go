/**
 * Author: Deean
 * Date: 2023-10-15 09:43
 * FileName: lib/Strings.go
 * Description:
 */

package lib

import (
	"strconv"
	"strings"
)

func Str2Array(s string) []int {
	ss := strings.Split(s[1:len(s)-1], ",")
	nums := []int{}
	for _, el := range ss {
		num, _ := strconv.Atoi(el)
		nums = append(nums, num)
	}
	return nums
}

func Str2Array2D(s string) [][]int {
	s = strings.ReplaceAll(s, "],[", "]#[")
	ss := strings.Split(s[1:len(s)-1], "#")
	grid := make([][]int, 0, len(ss))
	for _, el := range ss {
		grid = append(grid, Str2Array(el))
	}
	return grid
}
