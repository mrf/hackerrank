package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.ParseInt(scanner.Text(), 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(factorial(int(number)))

}

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	}
	return num
}
