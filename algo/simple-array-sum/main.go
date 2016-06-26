package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

var size int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var count int
	var data string
	var sum int
	count = 0
	for scanner.Scan() {
		count = count + 1
		line := scanner.Text()
		if count == 1 {
			//data_length = line
		} else if count == 2 {
			data = line
		}
		if line == "" {
			break
		}
	}
	var strings = strings.Split(data, " ")

	for i := 0; i < len(strings); i++ {
		converted_string, err := strconv.Atoi(strings[i])
		if err != nil {
			fmt.Println(err)
		}
		sum = converted_string + sum
	}

	fmt.Println(sum)
}
