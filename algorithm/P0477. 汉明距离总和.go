/**
 * Author: Deean
 * Date: 2022/12/10 17:20
 * FileName: algorithm/P0477. 汉明距离总和.go
 * Description:
 */

package main

import "fmt"

func totalHammingDistance(nums []int) int {
	n, total := len(nums), 0
	for i := 0; i <= 31; i++ {
		n1 := 0
		for _, v := range nums {
			n1 += v >> i & 1
		}
		total += n1 * (n - n1)
	}
	return total
}

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}
