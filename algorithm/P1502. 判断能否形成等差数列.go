/**
 * Author: Deean
 * Date: 2022-10-20 23:38
 * FileName: algorithm/P1502. 判断能否形成等差数列.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}
