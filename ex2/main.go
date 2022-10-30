// https://goplay.tools/snippet/Y4Ge1zvAohr

package main

import "fmt"

func stone(x int, s []int, f func([]int)) {
	gen(x, 0, s, nil, f)
}
func gen(x, sum int, s, tail []int, f func([]int)) {
	fmt.Println(s)
	if sum == x {
		f(tail)
	} else {
		for i := 0; sum < x; i++ {
			tail = append(tail, s[i])
			sum += s[i]
			gen(x, sum, s[i+1:], tail, f)
		}
	}
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	stone(4, s, func(v []int) {
		fmt.Println(v)
	})
}
