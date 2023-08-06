package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	real_n, _ := strconv.Atoi(n)

	for i := 1; i <= real_n; i++ {
		fmt.Print("ー")
	}

	fmt.Println()
}
