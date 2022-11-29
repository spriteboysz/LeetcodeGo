/**
 * Author: Deean
 * Date: 2022/11/29 22:42
 * FileName: algorithm/P0811. 子域名访问计数.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	hash := map[string]int{}
	for _, domain := range cpdomains {
		i := strings.IndexByte(domain, ' ')
		cnt, _ := strconv.Atoi(domain[:i])
		domain = domain[i+1:]
		hash[domain] += cnt

		for {
			i = strings.IndexByte(domain, '.')
			if i < 0 {
				break
			}
			domain = domain[i+1:]
			hash[domain] += cnt
		}
	}
	ans := make([]string, 0, len(hash))
	for s, c := range hash {
		ans = append(ans, strconv.Itoa(c)+" "+s)
	}
	return ans
}

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
