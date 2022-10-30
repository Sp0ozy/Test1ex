// https://goplay.tools/snippet/Y4Ge1zvAohr

package main

import (
	"fmt"
)

func stone(x int, s []int, f func([]int)) {
	gen(x, s, nil, f)
}
func gen(x int, s []int, tail []int, f func([]int)) {
	if x == 0 {
		f(tail)
	} else {
		for i := 0; i <= len(s)-1; i++ {
			if len(tail) == 0 || s[i] >= tail[len(tail)-1] {
				if len(s) > i+2 {
					gen(x-s[i], s[i+1:], append(tail, s[i]), f)
				}
			}
		}

	}
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	stone(17, s, func(v []int) {
		fmt.Println(v)
	})
}
