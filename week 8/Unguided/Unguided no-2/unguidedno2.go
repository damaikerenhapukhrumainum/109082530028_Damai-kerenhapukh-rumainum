package main

import (
	"fmt"
	"math"
)

func main() {
	var arr []int
	var n int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var nilai int
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scan(&nilai)
		arr = append(arr, nilai)
	}

	fmt.Println("Isi array:", arr)

	fmt.Println("Index ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}

	fmt.Println("Index genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}

	var x int
	fmt.Print("Masukkan kelipatan x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen kelipatan", x, ":")
	for _, v := range arr {
		if v%x == 0 {
			fmt.Println(v)
		}
	}

	var indexHapus int
	fmt.Print("Masukkan index yang akan dihapus: ")
	fmt.Scan(&indexHapus)

	arr = append(arr[:indexHapus], arr[indexHapus+1:]...)
	fmt.Println("Setelah dihapus:", arr)

	var total int
	for _, v := range arr {
		total += v
	}

	rata := float64(total) / float64(len(arr))
	fmt.Println("Rata-rata:", rata)

	var jumlahKuadrat float64
	for _, v := range arr {
		jumlahKuadrat += math.Pow(float64(v)-rata, 2)
	}

	stdDev := math.Sqrt(jumlahKuadrat / float64(len(arr)))
	fmt.Println("Standar deviasi:", stdDev)

	var cari int
	var frekuensi int

	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)

	for _, v := range arr {
		if v == cari {
			frekuensi++
		}
	}

	fmt.Println("Frekuensi", cari, "=", frekuensi)
}