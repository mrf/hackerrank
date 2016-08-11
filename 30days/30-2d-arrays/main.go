package main

import "fmt"
import "bufio"
import "os"

//import "strconv"
//import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	stuff := make([][]int, 6)
	for i := 0; i < 6; i++ {
		scanner.Scan()
		row := make([]int, 6)
		for r := 0; r < 6; r++ {
			// todo actually read each value in from line fmt.Println(scanner.text())
			row[r] = 66
		}

		stuff[i] = row
	}
	fmt.Println(stuff)
}
