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

	for _, arg := range os.Args[1:] {
		if arg == "END" {
			os.Exit(0)
		}
		num, _ := strconv.ParseInt(arg, 0, 64)
		fmt.Println(num)
	}
}
