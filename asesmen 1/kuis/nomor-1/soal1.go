package main
import "fmt"
const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 > m2 {
		fmt.Println("Selisih massa zat cair kiri dan massa zat cair kanan :", m1-m2)
	} else {
		fmt.Println("Selisih massa zat cair kiri dan massa zat cair kanan :", m2-m1)
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mJKiri, mJKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukan jari jari alas tabung: ")
	fmt.Scan(&r)
	fmt.Print("Masukan tinggi zat cair tabung kiri: ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukan massa jenis zat cair tabung kiri: ")
	fmt.Scan(&mJKiri)
	fmt.Print("Masukan tinggi zat cair tabung kanan: ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukan massa jenis zat cair tabunng kanan: ")
	fmt.Scan(&mJKanan)

	massaKiri = massa(r, tKiri, mJKiri)
	massaKanan = massa(r, tKanan, mJKanan)

	display(massaKiri, massaKanan)
}