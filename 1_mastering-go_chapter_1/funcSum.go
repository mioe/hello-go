package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("An error")
	k := 1

	var n float64
	var sum float64 = 0

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("None of the arguments is a float!")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}

	for i := 1; i < len(arguments); i++ {
		n, err = strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum = sum + n
		}
	}

	fmt.Println("Sum:", sum)
}
