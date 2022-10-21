/**
 * Author: Deean
 * Date: 2022-10-21 23:09
 * FileName: algorithm/P1394. 找出数组中的幸运数.go
 * Description:
 */

package main

import "fmt"

func findLucky(arr []int) int {
	hash := map[int]int{}
	for _, num := range arr {
		hash[num]++
	}
	lucky := []int{}
	for k, v := range hash {
		if k == v {
			lucky = append(lucky, k)
		}
	}
	if len(lucky) > 0 {
		maximum := 0
		for _, num := range lucky {
			if maximum < num {
				maximum = num
			}
		}
		return maximum
	} else {
		return -1
	}
}

func main() {
	arr := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(findLucky(arr))
}
