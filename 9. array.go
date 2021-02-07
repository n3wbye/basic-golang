package main

import "fmt"

func main() {

	// 1. Pengisian Elemen Array yang Melebihi Alokasi Awal

	var names [4]string

	names[0] = "awesome"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])

	// 2. Inisialisasi Nilai Awal Array

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// 3. Inisialisasi Nilai Array Dengan Gaya Vertikal

	fruits = [4]string{"guava", "orange", "manggo", "melon"}

	fruits = [4]string{
		"guava",
		"orange",
		"manggo",
		"melon",
	}

	// 4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen

	var numbers = [...]int{2, 3, 4, 5, 6}

	fmt.Println("data array \t: ", numbers)
	fmt.Println("jumlah elemen \t: ", len(numbers))

	// 5. Array Multidimensi

	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 2, 4}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// 5. Perulangan Elemen Array Menggunakan Keyword for

	fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	// 6. Perulangan Elemen Array Menggunakan Keyword for - range

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	// 7. Penggunaan Variabel Underscore _ Dalam for - range

	fruits = [4]string{"apple", "grape", "banana", "melon"}
	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// 8. Alokasi Elemen Array Menggunakan Keyword make

	var fruits3 = make([]string, 2)
	fruits3[0] = "apple"
	fruits3[1] = "manggo"

	fmt.Println(fruits3)

}
