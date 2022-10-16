/**
 * Author: Deean
 * Date: 2022-10-16 16:33
 * FileName: algorithm/P2103. 环和杆.go
 * Description:
 */

package main

import "fmt"

func countpoints2103(rings string) int {
	var rods [10]int
	for i := 0; i < len(rings); i += 2 {
		switch rings[i] {
		case 'B':
			rods[rings[i+1]-'0'] |= 1
		case 'R':
			rods[rings[i+1]-'0'] |= 2
		case 'G':
			rods[rings[i+1]-'0'] |= 4
		}
	}
	cnt := 0
	for _, rod := range rods {
		if rod == 7 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countpoints2103("B0R0G0R9R0B0G0"))
}
