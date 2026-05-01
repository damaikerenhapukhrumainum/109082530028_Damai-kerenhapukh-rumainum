package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan banyak ikan yang dijual dan banyak wadah: ")
	fmt.Scan(&x, &y)

	var berat [1000]float64
	var total [1000]float64

	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		wadah := i % y
		total[wadah] += berat[i]
	}

	var jumlahSemua float64

	fmt.Println("Total berat ikan pada setiap wadah:")
	for i := 0; i < y; i++ {
		fmt.Printf("Wadah ke-%d: %.2f\n", i+1, total[i])
		jumlahSemua += total[i]
	}

	rataRata := jumlahSemua / float64(y)
	fmt.Printf("Rata-rata berat ikan setiap wadah: %.2f\n", rataRata)
}