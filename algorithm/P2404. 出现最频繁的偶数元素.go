/**
 * Author: Deean
 * Date: 2022/11/20 21:33
 * FileName: algorithm/P2404. 出现最频繁的偶数元素.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func mostFrequentEven(nums []int) int {
	hash := map[int]int{}
	for _, num := range nums {
		if num%2 == 0 {
			hash[num]++
		}
	}
	maximum := -1
	for _, v := range hash {
		if v > maximum {
			maximum = v
		}
	}
	evens := []int{}
	for k, v := range hash {
		if v == maximum {
			evens = append(evens, k)
		}
	}
	if len(evens) == 0 {
		return -1
	}
	sort.Ints(evens)
	return evens[0]
}

func main() {
	fmt.Println(mostFrequentEven([]int{4, 4, 4, 9, 2, 4}))
}
