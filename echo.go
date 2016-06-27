// echo prints its command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	command := os.Args[0]
	fmt.Println("Command: " + command)
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i+1) + ": " + arg)
	}
}
