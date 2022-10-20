/**
 * Author: Deean
 * Date: 2022-10-20 23:46
 * FileName: sword/剑指 Offer 39. 数组中出现次数超过一半的数字.go
 * Description:
 */

package main

import "fmt"

func majorityElement(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		hash[num]++
		if hash[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(nums))
}
