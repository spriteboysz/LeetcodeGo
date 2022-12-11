/**
 * Author: Deean
 * Date: 2022/12/11 17:56
 * FileName: algorithm/P2240. 买钢笔和铅笔的方案数.go
 * Description:
 */

package main

import "fmt"

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	cnt := int64(0)
	for i := 0; i <= total/cost1; i++ {
		cnt += int64((total-cost1*i)/cost2) + 1
	}
	return cnt
}

func main() {
	fmt.Println(waysToBuyPensPencils(20, 10, 5))
}
