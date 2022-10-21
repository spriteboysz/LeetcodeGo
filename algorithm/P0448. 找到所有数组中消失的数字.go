/**
 * Author: Deean
 * Date: 2022-10-21 22:44
 * FileName: algorithm/P0448. 找到所有数组中消失的数字.go
 * Description:
 */

package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	var missing []int
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !hash[i] {
			missing = append(missing, i)
		}
	}
	return missing
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
