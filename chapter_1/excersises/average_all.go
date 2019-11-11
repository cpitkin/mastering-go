package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	var total, average float64

	for _, arg := range os.Args[1:] {
		num, _ := strconv.ParseFloat(arg, 64)
		total = total + num
		average = total / float64(len(os.Args[1:]))
	}

	fmt.Println(average)
}
