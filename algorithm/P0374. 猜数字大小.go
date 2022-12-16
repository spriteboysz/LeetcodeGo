/**
 * Author: Deean
 * Date: 2022/12/15 23:14
 * FileName: algorithm/P0374. 猜数字大小.go
 * Description:
 */

package main

import "fmt"

func guess(num int) int {
	if num > 6 {
		return -1
	} else if num < 6 {
		return 1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		if guess(mid) == 1 {
			left = mid + 1
		} else if guess(mid) == -1 {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	fmt.Println(guessNumber(10))
}
