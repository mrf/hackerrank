package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var m map[string]int
	m = make(map[string]int)
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < length; i++ {
		scanner.Scan()
		pieces := strings.Fields(scanner.Text())
		number, err := strconv.Atoi(pieces[1])
		if err != nil {
			fmt.Println(err)
		}

		m[pieces[0]] = number
	}

	// Search continues until it doesn't
	for scanner.Scan() {
		search := scanner.Text()
		// TODO if name matches print out name=value
		if m[search] > 0 {
			fmt.Print(search, "=", m[search])
		} else {
			fmt.Print("Not found")
		}
		fmt.Println()
	}
}
