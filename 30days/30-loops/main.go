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
		fmt.Println(err)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(integer, "x", i, "=", (integer * i))
	}
}
