package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// memanggil fungsi tanpa nilai balik
	var names = []string{"John", "Wick"}
	printMessage("Halooo", names)

	rand.Seed(time.Now().Unix())

	// memanggil fungsi tanpa nilai balik
	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Nilai random: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Nilai random: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Nilai random: ", randomValue)

	// memanggil funsi devideNumber()
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(6, 3)
}

// fungsi tanpa nilai balik
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan nilai balik
func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

// return untuk menghentikan proses dalam fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannont divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
