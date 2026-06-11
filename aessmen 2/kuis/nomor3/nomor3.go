package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func inputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("=== Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ===")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func provinsiTercepat(tumbuh TumbuhProv) int {
	idx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idx] {
			idx = i
		}
	}
	return idx
}

func indeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("=== Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ===")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil := int(float64(pop[i]) * (1 + tumbuh[i]))
			fmt.Println(prov[i], hasil)
		}
	}
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var namaCari string

	inputData(&prov, &pop, &tumbuh)

	fmt.Scan(&namaCari)

	idxCepat := provinsiTercepat(tumbuh)
	fmt.Println()
	fmt.Println("Provinsi dengan angka pertumbuhan tercepat :", prov[idxCepat])
	fmt.Println()

	idxCari := indeksProvinsi(prov, namaCari)
	fmt.Println("Data provinsi yang dicari :", namaCari)
	if idxCari == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println(prov[idxCari], pop[idxCari], tumbuh[idxCari])
	}

	fmt.Println()
	prediksi(prov, pop, tumbuh)
}