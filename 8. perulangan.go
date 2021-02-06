package main

import "fmt"

func main() {

	// 1. Perulangan Menggunakan Keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// 2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi
	var a = 0

	for a < 5 {
		fmt.Println("4ngka", a)
		a++
	}

	// 3. Penggunaan Keyword for Tanpa Argumen
	a = 0

	for {
		fmt.Println("Angk4", a)

		a++
		if a == 5 {
			break
		}
	}

	// 4. Penggunaan Keyword for - range

	/*
		Cara ke-4 adalah perulangan dengan menggunakan kombinasi keyword for dan range.
		Cara ini biasa digunakan untuk me-looping data bertipe array.
		Detailnya akan dibahas dalam bab selanjutnya (bab 14).
	*/

	// 5. Penggunaan Keyword break & continue

	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angkae", i)
	}

	// 6. Perulangan Bersarang

	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	// 7. Pemanfaatan Label Dalam Perulangan

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]\n")
		}
	}

}
