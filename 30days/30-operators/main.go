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
	fmt.Println(mealCost)
	scanner.Scan()
	tipPercent, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tipPercent)
	scanner.Scan()
	taxPercent, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(taxPercent)

	var tip float64
	tip = (mealCost * (tipPercent / 100))
	fmt.Println(tip)
	var tax float64
	tax = (mealCost * (taxPercent / 100))
	fmt.Println(tax)
}
