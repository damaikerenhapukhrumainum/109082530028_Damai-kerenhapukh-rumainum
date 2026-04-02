package main
import "fmt"
func main() {
	var beratGram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)

	kg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram > 500 {
			biayaSisa = sisaGram * 5
		} else if sisaGram > 0 {
			biayaSisa = sisaGram * 15
		}
	}

	total := biayaKg + biayaSisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Printf("Detail berat : %d kg + %d gram\n", kg, sisaGram)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", total)
}