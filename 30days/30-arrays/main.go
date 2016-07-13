package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	scanner.Scan()
	numbers := strings.Fields(scanner.Text())

	for i := length - 1; i > -1; i-- {
		fmt.Printf("%s ", numbers[i])
	}
}
