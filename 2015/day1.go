package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		fmt.Println("Failed to read input file.")
		return
	}
	defer file.Close()

	floor := 0
	current_index := 0
	basement_index := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		if scanner.Text() == "(" {
			floor += 1
		} else {
			floor -= 1
		}

		if basement_index == 0 && floor == -1 {
			basement_index = current_index + 1
		}

		current_index += 1
	}

	fmt.Printf("%d\n%d\n", floor, basement_index)
}
