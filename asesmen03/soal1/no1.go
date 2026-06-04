package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i

		for j := i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}

		T[i], T[idxMin] = T[idxMin], T[i]
	}
}

func median(T arrInt, n int) float64 {
	if n%2 == 1 {
		return float64(T[n/2])
	} else {
		tengah1 := T[n/2-1]
		tengah2 := T[n/2]
		return float64(tengah1+tengah2) / 2.0
	}
}

func main() {
	var data arrInt
	var n int
	var x int

	fmt.Println("Input data masukan :")

	for {
		fmt.Scan(&x)

		if x == -5313541 {
			break
		}

		if x == 0 {
			selectionSort(&data, n)
			fmt.Println("Median :")
			fmt.Println(median(data, n))
		} else {
			data[n] = x
			n++
		}
	}
}
