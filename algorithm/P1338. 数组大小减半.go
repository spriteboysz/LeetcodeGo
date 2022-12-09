/**
 * Author: Deean
 * Date: 2022/12/9 22:53
 * FileName: algorithm/P1338. 数组大小减半.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func minSetSize(arr []int) int {
	hash := map[int]int{}
	for _, num := range arr {
		hash[num]++
	}
	d := [][]int{}
	for k, v := range hash {
		d = append(d, []int{k, v})
	}
	sort.Slice(d, func(i, j int) bool {
		return d[i][1] > d[j][1]
	})
	c, l := 0, len(arr)/2
	for i, item := range d {
		c += item[1]
		if c > l {
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println(minSetSize([]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}))
}
