/**
 * Author: Deean
 * Date: 2022-10-16 23:30
 * FileName: algorithm/P1748. 唯一元素的和.go
 * Description:
 */

package main

import "fmt"

func sumOfUnique(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
	}
	sum := 0
	for k, v := range hash {
		if v == 1 {
			sum += k
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfUnique(nums))
}
