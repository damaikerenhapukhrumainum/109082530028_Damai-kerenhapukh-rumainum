package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

func insertionSort(t *tabPartai, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i - 1

		for j >= 0 && t[j].suara < temp.suara {
			t[j+1] = t[j]
			j--
		}

		t[j+1] = temp
	}
}

func main() {
	var p tabPartai
	var nPartai int
	var x int

	fmt.Println("Masukkan proses input suara :")

	for {
		fmt.Scan(&x)

		if x == -1 {
			break
		}

		idx := posisi(p, nPartai, x)

		if idx == -1 {
			p[nPartai].nama = x
			p[nPartai].suara = 1
			nPartai++
		} else {
			p[idx].suara++
		}
	}

	insertionSort(&p, nPartai)

	fmt.Println()
	fmt.Println("Hasil Perhitungan suara :")
	for i := 0; i < nPartai; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
}
