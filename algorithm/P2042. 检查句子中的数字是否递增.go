/**
 * Author: Deean
 * Date: 2022/11/12 23:01
 * FileName: algorithm/P2042. 检查句子中的数字是否递增.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	words := strings.Split(s, " ")
	nums := []int{}
	for _, word := range words {
		if word[0] >= '1' && word[0] <= '9' {
			num, _ := strconv.Atoi(word)
			nums = append(nums, num)
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areNumbersAscending("1 box has 3 blue 4 red 6 green and 12 yellow marbles"))
}
