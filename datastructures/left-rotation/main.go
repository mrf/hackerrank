package main

import "fmt"
import "bufio"
import "os"
import "strings"

import "strconv"

import "github.com/davecgh/go-spew/spew"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()
	fmt.Println("line1:", line1)
	scanner.Scan()

	line2 := scanner.Text()
	fmt.Println("Line2:", line2)

	control := strings.Fields(line1)

	length, _ := strconv.Atoi(control[0])
	rotations, _ := strconv.Atoi(control[1])
	//var numbers [length]int
	//rotations := strconv.Atoi(control[1])
	fmt.Printf("%T", length)
	fmt.Printf("%T", rotations)

	numbers := strings.Fields(line2)
	spew.Dump(numbers)
	/*
		for i < rotations {

		}
	*/

}
