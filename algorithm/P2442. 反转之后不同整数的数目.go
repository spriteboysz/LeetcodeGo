/**
 * Author: Deean
 * Date: 2022/11/29 23:39
 * FileName: algorithm/P2442. 反转之后不同整数的数目.go
 * Description:
 */

package main

import "fmt"

func countDistinctIntegers(nums []int) int {
	hash := map[int]bool{}
	for _, x := range nums {
		hash[x] = true
		rev := 0
		for ; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		hash[rev] = true
	}
	return len(hash)
}

func main() {
	fmt.Println(countDistinctIntegers([]int{1, 13, 10, 12, 31}))
}
