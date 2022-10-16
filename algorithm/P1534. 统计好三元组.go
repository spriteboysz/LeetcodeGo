/**
 * Author: Deean
 * Date: 2022-10-16 23:11
 * FileName: algorithm/P1534. 统计好三元组.go
 * Description:
 */

package main

import "fmt"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	cnt, n := 0, len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					cnt++
				}
			}
		}
	}
	return cnt
}

func main() {
	arr := []int{1, 1, 2, 2, 3}
	fmt.Println(countGoodTriplets(arr, 0, 0, 1))
}
