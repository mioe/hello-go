package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n float64
	var sum float64 = 0
	var err error = errors.New("An error")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()
		if t == "END" {
			fmt.Println("Goodbye")
			os.Exit(1)
		}

		n, err = strconv.ParseFloat(t, 64)
		if err == nil {
			sum = sum + n
			fmt.Println("> ", t, " | sum = ", sum)
		}
	}
}
