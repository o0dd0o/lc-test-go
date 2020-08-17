package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	qmap := map[string][]string{}
	for _, v := range tickets {
		qmap[v[0]] = append(qmap[v[0]], v[1])
	}
	for _, v := range qmap {
		sort.Strings(v)
	}
	re := []string{}
	var dfs func(string2 string)
	dfs = func(string2 string) {
		for len(qmap[string2]) > 0 {
			s := qmap[string2][0]
			qmap[string2] = qmap[string2][1:]
			dfs(s)
		}
		re = append([]string{string2}, re...)
	}
	dfs("JFK")
	return re
}
func main() {
	v1 := [][]string{
		{"MUC", "LHR"}, {"JFK", "MUC"},
		{"SFO", "SJC"}, {"LHR", "SFO"}}
	fmt.Printf("%v\n", findItinerary(v1))
	v2 := [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"},
		{"SFO", "ATL"}, {"ATL", "JFK"}}
	fmt.Printf("%v\n", findItinerary(v2))
}
