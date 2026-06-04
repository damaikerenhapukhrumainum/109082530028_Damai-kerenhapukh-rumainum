package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1001

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func sorting(data *arrPemain, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if data[j].gol > data[idxMax].gol ||
				(data[j].gol == data[idxMax].gol && data[j].assist > data[idxMax].assist) {
				idxMax = j
			}
		}
		data[i], data[idxMax] = data[idxMax], data[i]
	}
}

func main() {
	var data arrPemain
	var n int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan Data Input :")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		baris, _ := reader.ReadString('\n')
		baris = strings.TrimSpace(baris)

		parts := strings.Fields(baris)
		data[i].gol = 0
		data[i].assist = 0

		fmt.Sscanf(parts[len(parts)-2], "%d", &data[i].gol)
		fmt.Sscanf(parts[len(parts)-1], "%d", &data[i].assist)

		data[i].nama = strings.Join(parts[:len(parts)-2], " ")
	}

	sorting(&data, n)

	fmt.Println()
	fmt.Println("Hasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Println(data[i].nama, data[i].gol, data[i].assist)
	}
}
