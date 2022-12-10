/**
 * Author: Deean
 * Date: 2022/12/10 16:29
 * FileName: algorithm/P0869. 重新排序得到 2 的幂.go
 * Description:
 */

package main

import "fmt"

func reorderedPowerOf2(n int) bool {
	var target []int
	for i := 1; i <= 1e9; i *= 2 {
		target = append(target, i)
	}
	numbers1 := [10]int{}
	for ; n != 0; n /= 10 {
		numbers1[n%10]++
	}
	var numbers2 [10]int
Loop:
	for _, k := range target {
		numbers2 = [10]int{}
		for ; k != 0; k /= 10 {
			numbers2[k%10]++
		}
		for i := 0; i < 10; i++ {
			if numbers1[i] != numbers2[i] {
				continue Loop
			}
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(reorderedPowerOf2(1))
	fmt.Println(reorderedPowerOf2(10))
}
