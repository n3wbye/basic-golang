package main

import "fmt"

func main() {
	// 1. Tipe data numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -123123123

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// 2. Tipe data numerik desimal
	var decimalNumber = 2.26

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// 3. Tipe data bool (boolean)
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)

	// 4. Tipe data string
	var message string = "Hallo"
	fmt.Printf("message: %s\n", message)

	message = `Nama saya "John Wick"
	 Salam kenal.
	 Mari belajar golang
	`

	fmt.Printf("message: %s\n", message)

	// 5. Nilai nil & Zero Value

	/*
		nil bukan merupakan tipe data, melainkan sebuah nilai.
		Variabel yang isi nilainya nil berarti memiliki nilai kosong.
	*/
	/*
		Zero value dari string adalah "" (string kosong).
		Zero value dari bool adalah false.
		Zero value dari tipe numerik non-desimal adalah 0.
		Zero value dari tipe numerik desimal adalah 0.0.
	*/
}
