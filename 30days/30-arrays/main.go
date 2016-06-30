package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	scanner.Scan()
	numbers := scanner.Text()
	fmt.Println(length)
	arr := [length]int{1, 3, 4, 5}
	fmt.Println(arr)
	/*
		for number in arr {
			fmt.Println(number)
		}
	*/
}
