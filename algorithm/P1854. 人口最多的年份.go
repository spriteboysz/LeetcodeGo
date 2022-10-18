/**
 * Author: Deean
 * Date: 2022-10-18 23:58
 * FileName: algorithm/P1854. 人口最多的年份.go
 * Description:
 */

package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	offset := 1950
	growth := make([]int, 2050-offset+1)
	for _, v := range logs {
		growth[v[0]-offset]++
		growth[v[1]-offset]--
	}

	sum, max, maxYear := 0, 0, 0
	for i, v := range growth {
		sum += v
		if sum > max {
			max = sum
			maxYear = i
		}
	}
	return maxYear + offset
}

func main() {
	logs := [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}
	fmt.Println(maximumPopulation(logs))
}
