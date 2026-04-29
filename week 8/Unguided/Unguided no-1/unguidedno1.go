package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	pusat  Titik
	radius float64
}

func jarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= c.radius
}

func main() {
	var l1, l2 Lingkaran
	var p Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	dalam1 := didalam(l1, p)
	dalam2 := didalam(l2, p)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}