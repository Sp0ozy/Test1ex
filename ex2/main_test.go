package main

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestStone(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		x := rand.Intn(30)
		m := make(map[string]int)
		stone(x, []int{1, 11, 3, 13, 4, 6, 7, 8, 9, 10, 2, 12, 5, 15, 14, 16, 17, 18, 19, 20}, func(v []int) {
			var res string
			sort.Ints(v)
			for _, x := range v {
				res += strconv.Itoa(x)
			}
			if m[res] > 0 {
				t.Errorf("%v was already shown %v", v, x)
			} else {
				m[res]++
			}
			var sum int
			for _, x := range v {
				sum += x
			}
			if sum != x {
				t.Errorf("sum of %v is %v, want %v", v, sum, x)
			}
		})
	}
}
