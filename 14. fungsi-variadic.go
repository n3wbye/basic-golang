package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)

	// pengisian paramter fungsi variadic menggunakan data slice
	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	avg = calculate(numbers...)
	fmt.Printf("Rata-rata : %.2f\n", avg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, numbers := range numbers {
		total += numbers
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
