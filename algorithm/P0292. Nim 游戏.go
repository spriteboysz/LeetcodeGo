/**
 * Author: Deean
 * Date: 2022-10-20 23:20
 * FileName: algorithm/P0292. Nim 游戏.go
 * Description:
 */

package main

import "fmt"

func canWinNim(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(canWinNim(4))
	fmt.Println(canWinNim(1))
	fmt.Println(canWinNim(2))
}
