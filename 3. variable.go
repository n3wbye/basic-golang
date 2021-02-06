package main

import "fmt"

func main() {
	// 1. Deklarasi variable beserta tipe data
	var firstName string = "Bae"

	var lastName string
	lastName = "Suzy"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// 2. Deklarasi variable tanpa tipe data
	firstName2 := "Alfa"
	lastName2 := "mart"

	fmt.Printf("Ada %s%s!\n", firstName2, lastName2)

	// *perlu diketahu, deklarasi menggunakan := hanya bisa dilakukan di dalam blok fungsi.

	// 3. Deklarasi multi variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// 4. Variabel underscore _ (tidak dipakai)
	_ = "belajar golang"
	name, _ := "john", "wick"

	// 5. Deklarasi variable menggunakan keywork new (pointer)
	name := new(string)

	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""

	// 6. Deklarasi variable menggunakan keyword make

	/*
		- channel
		- slice
		- map
	*/

}
