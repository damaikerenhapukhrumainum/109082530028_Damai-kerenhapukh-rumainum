package main

import "fmt"

func main() {
	var x int
	var totalMasuk, suaraSah int
	suara := make(map[int]int)

	for {
		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}

		if x == 0 {
			break
		}

		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}