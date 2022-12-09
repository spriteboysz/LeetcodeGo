/**
 * Author: Deean
 * Date: 2022/12/9 22:40
 * FileName: algorithm/P0973. 最接近原点的 K 个点.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

func main() {
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
