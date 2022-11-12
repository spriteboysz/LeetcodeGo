/**
 * Author: Deean
 * Date: 2022/11/12 22:49
 * FileName: algorithm/P2248. 多个数组求交集.go
 * Description:
 */

package main

import "fmt"

func intersection2(nums [][]int) []int {
	numbers := make([]int, 1001)
	for _, num := range nums {
		for _, n := range num {
			numbers[n] += 1
		}
	}
	section := []int{}
	for i := 1; i <= 1000; i++ {
		if numbers[i] == len(nums) {
			section = append(section, i)
		}
	}
	return section
}

func main() {
	fmt.Println(intersection2([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
}
