package main

import "fmt"
import "bufio"
import "os"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Skip number of strings
	for scanner.Scan() {
		myString := scanner.Text()
		var even []rune
		var odd []rune
		for i, r := range myString {
			if i%2 == 0 {
				even = append(even, r)
			} else {
				odd = append(odd, r)
			}
		}

		for _, character := range even {
			fmt.Printf("%c", character)
		}
		fmt.Printf(" ")
		for _, character := range odd {
			fmt.Printf("%c", character)
		}
		fmt.Println()
	}
}
