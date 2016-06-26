package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var matrix_length int
	var line int
	var line_1 string
	scanner.Scan()
	line_1 = scanner.Text()
	fmt.Println(line_1)
	matrix_length, err := strconv.Atoi(line_1)
	if err != nil {
		// If not able to convert to number should just bail here
		fmt.Println(err)
	}
	matrix := make([][][][]int, matrix_length)
	fmt.Println(matrix)
	for line < matrix_length {
		scanner.Scan()
		// TODO assign array to matrix
		converted, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// If not able to convert to number should just bail here
			fmt.Println(err)
		}
		matrix[line] = [1,314,1234]
		fmt.Println(scanner.Text())
		line = line + 1
	}
	fmt.Println(matrix)
	/*
		for scanner.Scan() {
			} else {
				fmt.Println(iteration)
				// TODO something like this but this aint working
				//matrix[iteration] = scanner.Text()
			}
				if iteration > matrix_length {
					fmt.Println("past bounds exiting")
					break
				}
			iteration = iteration + 1
		}
	*/

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
