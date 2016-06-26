package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	integer, err := strconv.Atoi(scanner.Text())
	if err != nil {
		// If not able to convert to number should just bail here
		fmt.Println(err)
	}
	if (integer % 2) == 1 {
		fmt.Println("Weird")
	}
	if (integer % 2) == 0 {
		if (integer >= 2) && (integer <= 5) {
			fmt.Println("Not Weird")
		}
		if (integer >= 6) && (integer <= 20) {
			fmt.Println("Weird")
		}
		if integer > 20 {
			fmt.Println("Not Weird")
		}
	}
}
