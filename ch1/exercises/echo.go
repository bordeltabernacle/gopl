// echo prints its command line arguments
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	command := os.Args[0]
	fmt.Printf("Command: %s\n\n", command)
	for i, arg := range os.Args[1:] {
		fmt.Println(strconv.Itoa(i+1) + ": " + arg)
	}
	fmt.Printf("\n%.10fs elapsed\n", time.Since(start).Seconds())
}
