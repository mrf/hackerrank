package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	mealCost, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}
	scanner.Scan()
	tipPercent, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}
	scanner.Scan()
	taxPercent, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}

	var tip float64
	tip = (mealCost * (tipPercent / 100))
	var tax float64
	tax = (mealCost * (taxPercent / 100))
	var totalCost float64
	var intTotalCost int
	totalCost = mealCost + tip + tax
	// Trick to round up no matter what
	intTotalCost = int(totalCost + .5)
	fmt.Println("The total meal cost is", intTotalCost, "dollars.")
}
