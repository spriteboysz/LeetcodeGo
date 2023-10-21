/**
 * Author: Deean
 * Date: 2023-10-19 23:31
 * FileName: LCR/LCR 177. 撞色搭配.go
 * Description:
 */

package main

import "fmt"

func sockCollocation(sockets []int) []int {
	hash := map[int]int{}
	for _, socket := range sockets {
		hash[socket]++
	}
	single := []int{}
	for _, socket := range sockets {
		if hash[socket] == 1 {
			single = append(single, socket)
		}
	}
	return single
}

func main() {
	fmt.Println(sockCollocation([]int{1, 2, 4, 1, 4, 3, 12, 3}))
}
