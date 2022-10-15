/**
 * Author: Deean
 * Date: 2022-10-15 21:26
 * FileName: algorithm/P1313. 解压缩编码列表.go
 * Description:
 */

package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var decompress []int
	for i := 1; i < len(nums); i += 2 {
		for j := 0; j < nums[i-1]; j++ {
			decompress = append(decompress, nums[i])
		}
	}
	return decompress
}

func main() {
	nums := []int{1, 1, 2, 3}
	fmt.Println(decompressRLElist(nums))
}
