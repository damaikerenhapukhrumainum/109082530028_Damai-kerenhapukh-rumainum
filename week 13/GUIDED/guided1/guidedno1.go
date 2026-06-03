package main

import "fmt"

func SelectionSortArray(arr *[8]int) {
	var idxMin int

	for i := 0; i < len(arr)-1; i++ {
		idxMin = i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}

		if idxMin != i {
			arr[i], arr[idxMin] = arr[idxMin], arr[i]
		}
	}
}

func main() {
	var arrAngka [8]int

	for i := 0; i < len(arrAngka); i++ {
		fmt.Printf("Masukkan data angka indeks ke-%d: ", i)
		fmt.Scan(&arrAngka[i])
	}

	fmt.Println("\n=== SEBELUM SORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}

	SelectionSortArray(&arrAngka)

	fmt.Println("\n\n=== SETELAH SORTING ===")
	for i := 0; i < len(arrAngka); i++ {
		fmt.Print(arrAngka[i], " ")
	}
	fmt.Println()
}