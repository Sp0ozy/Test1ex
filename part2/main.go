package main

import "fmt"

func partition(s []int) (i int) {
	l, r := 0, len(s)-1
	p := s[r/2]
	for {
		for s[l] < p {
			l++
		}
		for s[r] > p {
			r--
		}
		if l >= r {
			return r
		}
		s[l], s[r] = s[r], s[l]
		l++
		r--
		fmt.Println(s)
	}
}
func sort(s []int) {
	if len(s) < 2 {
		return
	}
	p := partition(s) + 1
	sort(s[:p])
	sort(s[p:])
}
func find(s []int, i int) int {
	if s == nil || len(s)-1 < i {
		return 0
	}
	sort(s)
	return s[i]
}
func main() {
	s := []int{1, 2, 54, 101, 4, 1}
	find(s, 5)
}
