/**
 * Author: Deean
 * Date: 2022-10-12 23:31
 * FileName: algorithm/P1672. 最富有客户的资产总量.go
 * Description:
 */

package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	maximum := 0
	for _, account := range accounts {
		cur := 0
		for _, item := range account {
			cur += item
		}
		if cur > maximum {
			maximum = cur
		}
	}
	return maximum
}

func main() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts))
}
