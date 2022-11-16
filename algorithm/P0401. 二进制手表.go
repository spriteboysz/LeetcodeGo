/**
 * Author: Deean
 * Date: 2022/11/16 23:25
 * FileName: algorithm/P0401. 二进制手表.go
 * Description:
 */

package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	times := []string{}
	for i := 0; i < 12*60; i++ {
		hh, mm := i/60, i%60
		if bits.OnesCount(uint(hh))+bits.OnesCount(uint(mm)) == turnedOn {
			times = append(times, fmt.Sprintf("%d:%02d", hh, mm))
		}
	}
	return times
}

func main() {
	fmt.Println(readBinaryWatch(1))
}
