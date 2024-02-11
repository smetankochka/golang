package main

import (
	"errors"
	"fmt"
)

type Grasshopper struct {
	place  int
	target int
}

type Jumper interface {
	WhereAmI() int      // выводит текущее положение кузнечика на линейке
	Jump() (int, error) // кузнечик прыгает к зерну. Выводит новое положение кузнечика, или ошибку, если он уже ест зерно
}

func (g *Grasshopper) Jump() (int, error) {
	if g.place == g.target {
		return g.place, errors.New("уже сьел зерно")
	}
	if g.place < g.target {
		if g.target-g.place <= 5 {
			g.place = g.target
		} else {
			g.place += 5
		}
	} else {
		if g.place-g.target <= 5 {
			g.place = g.target
		} else {
			g.place -= 5
		}
	}
	return g.place, nil
}

func (g *Grasshopper) WhereAmI() int {
	return g.place
}

func PlaceJumper(place, target int) Jumper {
	return &Grasshopper{
		place:  place,
		target: target,
	}
}

func main() {
	place := 0
	target := 3
	g := PlaceJumper(place, target)
	fmt.Println(g.WhereAmI())
	for {
		currPlace, err := g.Jump()
		if err != nil {
			break
		}
		fmt.Println(currPlace)
	}
}
