package main

import "fmt"

func partition(s []int) (i int) {
	l, r := 0, len(s)-1
	p := s[r/2]
	for {
		fmt.Println(s)
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

	}
}
func qsort(s []int, n int) int {
	if len(s) <= 2 {
		fmt.Println(s)
		return s[len(s)-1]
	}

	p := partition(s) + 1

	if n >= p {
		fmt.Println(p, n, "bigger")
		return qsort(s[p:], len(s)-p-1)
	} else {
		fmt.Println(p, n, "smaller")
		return qsort(s[:p], n)
	}
}
func main() {
	s := []int{1, 2, 54, 101, 4, 1, 5}
	fmt.Println(qsort(s, 0))
}
