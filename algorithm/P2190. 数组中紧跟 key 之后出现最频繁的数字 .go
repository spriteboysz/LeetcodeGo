/**
 * Author: Deean
 * Date: 2022/11/20 17:05
 * FileName: algorithm/P2190. 数组中紧跟 key 之后出现最频繁的数字 .go
 * Description:
 */

package main

import "fmt"

func mostFrequent(nums []int, key int) int {
	hash := map[int]int{}
	for i, num := range nums {
		if i != len(nums)-1 && num == key {
			hash[nums[i+1]]++
		}
	}
	maximum, target := 0, -1

	for k, v := range hash {
		if v > maximum {
			maximum = v
			target = k
		}
	}
	return target
}

func main() {
	fmt.Println(mostFrequent([]int{2, 2, 2, 2, 3}, 2))
}
