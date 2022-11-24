/**
 * Author: Deean
 * Date: 2022/11/24 23:10
 * FileName: algorithm/P1629. 按键持续时间最长的键.go
 * Description:
 */

package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans, m := keysPressed[0], releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if diff := releaseTimes[i] - releaseTimes[i-1]; diff > m {
			ans, m = keysPressed[i], diff
		} else if diff == m && keysPressed[i] > ans {
			ans = keysPressed[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(slowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
	fmt.Println(slowestKey([]int{1, 2}, "ba"))
}
