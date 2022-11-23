/**
 * Author: Deean
 * Date: 2022/11/23 22:47
 * FileName: algorithm/P2437. 有效时间的数目.go
 * Description:
 */

package main

import "fmt"

func countTime(time string) int {
	cnt := 0
	for i := 0; i < 24*60; i++ {
		hh, mm := i/60, i%60
		cur := fmt.Sprintf("%02d:%02d", hh, mm)
		flag := true
		for j := 0; j < 5; j++ {
			if time[j] != '?' && cur[j] != time[j] {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countTime("0?:0?"))
}
