package main

import "fmt"
import "bufio"
import "os"
import "strconv"

//import "strings"

var size int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var count int
	var data_length string
	var data string
	count = 0
	for scanner.Scan() {
		count = count + 1
		line := scanner.Text()
		if count == 1 {
			data_length = line
		} else if count == 2 {
			data = line
		}
		if line == "" {
			break
		}
	}
	fmt.Printf("Data lenght: %s \n", data_length)
	fmt.Printf("Data: %s \n", data)
	converted, err := strconv.Atoi(data_length)
	if err != nil {
		fmt.Printf("Converted: %i \n", converted)
		//var sums [converted]int
	} else {
		fmt.Println(err)
	}
	//var strings = strings.Split(data, " ")
}
