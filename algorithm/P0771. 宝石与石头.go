/**
 * Author: Deean
 * Date: 2022-10-13 23:19
 * FileName: algorithm/P0771. 宝石与石头.go
 * Description:
 */

package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	cnt := 0
	for _, stone := range stones {
		for _, jewel := range jewels {
			if jewel == stone {
				cnt++
				break
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
