package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var iteration int
	var matrix_length int
	iteration = 1
	var line_1 string

	for scanner.Scan() {
		if iteration == 1 {
			line_1 = scanner.Text()
			matrix_length, err := strconv.Atoi(line_1)
			if err != nil {
				// if not a number should just bail
				fmt.Println(err)
				break
			}
			matrix := make([][]int, matrix_length)
			fmt.Println(matrix)
		}
		if iteration > matrix_length {
			fmt.Println("past bounds exiting")
			break
		}
		iteration = iteration + 1
	}

	// todo three times for each array
	for i := 0; i < matrix_length; i++ {
		//	converted_string, err := strconv.Atoi(strings[i])
		/*
			if err != nil {
				fmt.Println(err)
			}
			sum = converted_string + sum
		*/
	}

	//fmt.Println(sum)
}
