package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	counts := make(map[string]int64)

	for _, s := range cpdomains {

		cpinfo := strings.Fields(s)

		count, _ := strconv.ParseInt(cpinfo[0], 0, 64)

		frags := strings.Split(cpinfo[1], ".")

		cur := ""
		for i := len(frags) - 1; i >= 0; i-- {

			// do not append period to top-level domain

			top := i == len(frags)-1

			if top {

				cur = frags[i] + cur

			} else {

				cur = frags[i] + "." + cur

			}

			if val, ok := counts[cur]; ok {

				counts[cur] = val + count

			} else {
				fmt.Println(cur)

				counts[cur] = 0 + count

			}

		}

	}

	var ans []string

	for dom, val := range counts {
		ans = append(ans, strconv.FormatInt(val, 10)+" "+dom)
	}

	return ans

}
