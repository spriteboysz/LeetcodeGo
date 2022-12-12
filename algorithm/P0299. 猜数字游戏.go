/**
 * Author: Deean
 * Date: 2022/12/11 23:00
 * FileName: algorithm/P0299. 猜数字游戏.go
 * Description:
 */

package main

import "fmt"

func getHint(secret string, guess string) string {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	bulls, cows := 0, 0
	cntS, cntG := [10]int{}, [10]int{}
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntS[secret[i]-'0']++
			cntG[guess[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntS[i], cntG[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func main() {
	fmt.Println(getHint("1123", "0111"))
}
