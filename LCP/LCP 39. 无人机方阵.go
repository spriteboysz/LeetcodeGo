/**
 * Author: Deean
 * Date: 2022/12/1 23:16
 * FileName: LCP/LCP 39. 无人机方阵.go
 * Description:
 */

package main

import "fmt"

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	hash := map[int]int{}
	m, n := len(source), len(source[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			hash[source[i][j]]++
			hash[target[i][j]]--
		}
	}
	cnt := 0
	for _, v := range hash {
		if v < 0 {
			v = -v
		}
		cnt += v
	}
	return cnt / 2
}

func main() {
	fmt.Println(minimumSwitchingTimes([][]int{{1, 3}, {5, 4}}, [][]int{{3, 1}, {6, 5}}))
}
