package main

import "fmt"

type suhu float64
func CelciusToReamur(celcius suhu) suhu {
	return (4.0 / 5.0) * celcius
}

func CelciusToFahrenheit(celcius suhu) suhu {
	return (9.0/5.0)*celcius + 32
}
func CelciusToKelvin(celcius suhu) suhu {
	return celcius + 273.15
}

func main() {
	var c suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&c)

	reamur := CelciusToReamur(c)
	fahrenheit := CelciusToFahrenheit(c)
	kelvin := CelciusToKelvin(c)

	fmt.Printf("\n%.2f celcius = %.2f reamur\n", c, reamur)
	fmt.Printf("%.2f celcius = %.2f fahrenheit\n", c, fahrenheit)
	fmt.Printf("%.2f celcius = %.2f kelvin\n", c, kelvin)
}