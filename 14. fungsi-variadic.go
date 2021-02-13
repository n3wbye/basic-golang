package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pengisian paramter fungsi variadic menggunakan data slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate(numbers...)
	fmt.Printf("Rata-rata : %.2f\n", avg)

	yourHobbies("Asep", "Ngoding", "Hacking", "Dancing")
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, numbers := range numbers {
		total += numbers
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

// fungsi dengan parameter biasa & variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
