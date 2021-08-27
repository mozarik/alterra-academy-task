package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	countedItems := CountItems([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})

	p := make(PairList, len(countedItems))
	i := 0
	for k, v := range countedItems {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}
}

// Count the number of items in the slice and make a map of the items and their counts
// and return the map
func CountItems(items []string) map[string]int {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}
