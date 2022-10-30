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
func find(s []int, n int) int {
	if len(s)-1 < n || s == nil {
		return 0
	}
	if len(s) < 2 {
		return s[0]
	}
	p := partition(s) + 1
	if n >= p {
		fmt.Println(p, n, "bigger")
		return find(s[p:], n-p)
	} else {
		fmt.Println(p, n, "smaller")
		return find(s[:p], n)
	}
}
func main() {
	s := []int{1, 2, 54, 101, 4, 1}
	fmt.Println(find(s, 23))
}

//O(log(n))
