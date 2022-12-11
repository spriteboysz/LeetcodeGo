/**
 * Author: Deean
 * Date: 2022/12/11 15:48
 * FileName: algorithm/P1268. 搜索推荐系统.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	suggested := [][]string{}
	for i := 0; i < len(searchWord); i++ {
		target, result := searchWord[0:i+1], []string{}
		for _, product := range products {
			if len(product) >= len(target) && target == product[0:i+1] {
				result = append(result, product)
			}
			if len(result) == 3 {
				break
			}
		}
		suggested = append(suggested, result)
	}
	return suggested
}

func main() {
	fmt.Println(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
}
