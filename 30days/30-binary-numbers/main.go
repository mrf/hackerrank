package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "reflect"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.ParseInt(scanner.Text(), 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	var binary = strconv.FormatInt(number, 2)
	fmt.Println(binary)
	fmt.Println(reflect.TypeOf(binary))
	for i := 0; i < len(binary); i++ {
		if binary[i] == 49 {
			// TODO has to be a better way to do this
			if i-1 > 1 && binary[i-1] == 49 {
				fmt.Println("two in a row")
			}
		}
	}
}
