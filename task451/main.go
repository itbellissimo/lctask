package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "tree"

	fmt.Println(s)
	res := frequencySort(s)
	fmt.Println(res)
}

type freq struct {
	r   rune
	cnt int
}
type freqList []freq

func (l freqList) Len() int           { return len(l) }
func (l freqList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l freqList) Less(i, j int) bool { return l[i].cnt > l[j].cnt }

func frequencySort(s string) string {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}

	freqs := make(freqList, 0, len(m))
	for r, c := range m {
		freqs = append(freqs, freq{
			r:   r,
			cnt: c,
		})
	}

	sort.Sort(freqs)

	sb := strings.Builder{}
	for _, freq := range freqs {
		sb.WriteString(strings.Repeat(string(freq.r), freq.cnt))
	}

	return sb.String()
}
