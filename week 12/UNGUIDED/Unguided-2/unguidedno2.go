package main

import "fmt"

func main() {
	var x int
	var totalMasuk, suaraSah int
	suara := make(map[int]int)

	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}

		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			suaraSah++
		}
	}

	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}