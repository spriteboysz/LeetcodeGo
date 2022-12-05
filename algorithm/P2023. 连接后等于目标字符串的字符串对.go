/**
 * Author: Deean
 * Date: 2022/12/4 22:18
 * FileName: algorithm/P2023. 连接后等于目标字符串的字符串对.go
 * Description:
 */

package main

import "fmt"

func numOfPairs(nums []string, target string) int {
	n, cnt := len(nums), 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && len(nums[i])+len(nums[j]) == len(target) && nums[i]+nums[j] == target {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
}
