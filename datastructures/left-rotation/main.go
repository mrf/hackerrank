package main

import "fmt"
import "bufio"
import "os"
import "strings"

//import "strconv"

//`import "github.com/davecgh/go-spew/spew"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()
	fmt.Println("line1:", line1)
	scanner.Scan()

	line2 := scanner.Text()
	fmt.Println("Line2:", line2)

	control := strings.Split(line1, " ")

	length := control[0]
	fmt.Println(length)
	rotations := control[1]
	//rotations := strconv.Atoi(control[1])
	fmt.Printf("%T", rotations)
	/*
		for i < rotations {

		}
	*/

}
