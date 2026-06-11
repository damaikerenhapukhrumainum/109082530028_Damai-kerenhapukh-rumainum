package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func inputData(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].Nama, &T[i].Nilai)
	}
}

func cariNilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].Nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	max := -1

	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			if T[i].Nilai > max {
				max = T[i].Nilai
			}
		}
	}

	return max
}

func main() {
	var data arrayMahasiswa
	var N int
	var nimCari string

	inputData(&data, &N)

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(data, N, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(data, N, nimCari)

	fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", nilaiPertama)
	fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", nilaiTerbesar)
}