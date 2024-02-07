package main

import (
	"fmt"
)

type ogran struct {
	s  bool
	o1 int
	o2 int
}

type offset struct {
	ryad  ogran
	mesto ogran
}

func main() {
	var t, n, m, k, p, cur, si, ri, ci, q int
	fmt.Scan(&t)

	for t > 0 {
		t--
		fmt.Scan(&n, &m, &k, &p)
		count := k
		mas := make([][]struct {
			first  int
			second bool
		}, n)
		for i := range mas {
			mas[i] = make([]struct {
				first  int
				second bool
			}, m)
			for j := range mas[i] {
				mas[i][j] = struct {
					first  int
					second bool
				}{-1, false}
			}
		}

		og := make([]offset, k)
		used := make([]bool, k)
		var ty byte
		for i := 0; i < p; i++ {
			fmt.Scan(&ty, &cur)
			if ty == 'R' {
				fmt.Scan(&og[cur-1].ryad.o1, &og[cur-1].ryad.o2)
				og[cur-1].ryad.o1--
				og[cur-1].ryad.o2--
				og[cur-1].ryad.s = true
			} else {
				fmt.Scan(&og[cur-1].mesto.o1, &og[cur-1].mesto.o2)
				og[cur-1].mesto.o1--
				og[cur-1].mesto.o2--
				og[cur-1].mesto.s = true
			}
		}
		fmt.Scan(&q)
		for i := 0; i < q; i++ {
			var trash string
			fmt.Scan(&trash, &si, &ri, &ci)
			si--
			ri--
			ci--
			if !used[si] && mas[ri][ci].first == -1 {
				used[si] = true
				mas[ri][ci].first = si
				if og[si].ryad.s && (og[si].ryad.o1 > ri && og[si].ryad.o2 < ri) {
					mas[ri][ci].second = true
				} else if og[si].mesto.s && (og[si].mesto.o1 > ci && og[si].mesto.o2 < ci) {
					mas[ri][ci].second = true
				} else {
					count--
				}
			} else if used[si] {
				if mas[ri][ci].second && ((og[si].ryad.s && (og[si].ryad.o1 > ri && og[si].ryad.o2 < ri)) || (og[si].mesto.s && (og[si].mesto.o1 > ci || og[si].mesto.o2 < ci))) {
					used[si] = true
					used[mas[ri][ci].first] = false
					mas[ri][ci].first = si
					count--
				}
			}
		}
		fmt.Println(count)
	}
}
