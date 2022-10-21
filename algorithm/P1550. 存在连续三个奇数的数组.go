/**
 * Author: Deean
 * Date: 2022-10-21 23:19
 * FileName: algorithm/P1550. 存在连续三个奇数的数组.go
 * Description:
 */

package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	for i := 2; i < len(arr); i++ {
		if arr[i-2]%2 == 1 && arr[i-1]%2 == 1 && arr[i]%2 == 1 {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	fmt.Println(threeConsecutiveOdds(arr))
	arr = []int{1, 2, 3}
	fmt.Println(threeConsecutiveOdds(arr))
	arr = []int{1, 3, 23}
	fmt.Println(threeConsecutiveOdds(arr))
}
