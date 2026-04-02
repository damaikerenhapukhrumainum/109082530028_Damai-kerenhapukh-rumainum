package main
import "fmt"

func hitungBiaya(jenis string, masuk int, keluar int) int {
	jamParkir := keluar - masuk
	biaya := 0
	for i := 0; i < jamParkir; i++ {
	jamSekarang := (masuk + i) % 24
	if jenis == "motor" {
	if jamSekarang < 17 {
	biaya += 4000
	} else {
	biaya += 5000
	}
	} else if jenis == "mobil" {
	if jamSekarang < 17 {
	biaya += 6000
	} else {
	biaya += 7000
	}
	}
	}
	return biaya
	}
func main() {
	var jenis string
	var masuk, keluar int
	totalPendapatan := 0
	kendaraan := 1
	fmt.Println("==== Rekap Tarif Parkir Cafe Per Hari ====\n")
	for {
		fmt.Printf("*Kendaraan %d\n", kendaraan)
		fmt.Print("Kendaraan apa? (motor/mobil/- untuk selesai): ")
		fmt.Scan(&jenis)
		if jenis == "-" {
		break
		}
	fmt.Print("Masukkan Jam Masuk Kendaraan (0-24): ")
	fmt.Scan(&masuk)
	fmt.Print("Masukkan Jam Keluar Kendaraan (0-24): ")
	fmt.Scan(&keluar)
	biaya := hitungBiaya(jenis, masuk, keluar)
	fmt.Printf("Biaya parkir %s %d: %d\n", jenis, kendaraan, biaya)
	fmt.Println("===========================================")
	totalPendapatan += biaya
	kendaraan++
	}
	fmt.Printf("\n*** Total Pendapatan Parkir Hari Ini Adalah %d ***\n", totalPendapatan)
	}