/**
 * Author: Deean
 * Date: 2022/12/12 23:08
 * FileName: algorithm/P0692. 前K个高频单词.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func topKFrequent(words []string, k int) []string {
	hash := map[string]int{}
	for _, word := range words {
		hash[word]++
	}
	uniqueWords := make([]string, 0, len(hash))
	for w := range hash {
		uniqueWords = append(uniqueWords, w)
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		s, t := uniqueWords[i], uniqueWords[j]
		return hash[s] > hash[t] || hash[s] == hash[t] && s < t
	})
	return uniqueWords[:k]
}

func main() {
	fmt.Println(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4))
}
