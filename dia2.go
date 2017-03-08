package main

import (
	"fmt"
	"math"
)

func main() {
	totalCost(12.0, 20, 8)
}

func totalCost(mealCost float64, tipPercent int, taxPercent int) {
	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100

	total := mealCost + tip + tax

	t := math.Floor(total)
	fmt.Printf("The total meal cost is %v dollars.\n", t)

}
