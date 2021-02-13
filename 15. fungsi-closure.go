package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 9, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("Data: %v\nmax : %v\nmin : %v\n", numbers, max, min)
}
