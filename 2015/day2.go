package main

import (
	"fmt"
	"io"
	"os"
)

func min3(a int, b int, c int) int {
	var min int
	if a < b {
		min = a
	} else {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}

func surface_area(length int, width int, height int) int {
	return 2 * length * width + 2 * width * height + 2 * height * length
}

func smallest_side_area(length int, width int, height int) int {
	return min3(length * width, width * height, height * length)
}

func smallest_perimeter(length int, width int, height int) int {
	return min3(2 * length + 2 * width, 2 * width + 2 * height, 2 * height + 2 * length)
}

func volume(length int, width int, height int) int {
	return length * width * height
}

func main() {
    file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Failed to read input file.")
		return
	}
	defer file.Close()

	var length, width, height int
	wrapping_paper := 0
	ribbon := 0

	for {
		_, err := fmt.Fscanf(file, "%dx%dx%d", &length, &width, &height)
		if err != nil {
			if err == io.EOF {
				break;
			}
			fmt.Println(err)
			return
		}

		wrapping_paper += surface_area(length, width, height) + smallest_side_area(length, width, height)
		ribbon += smallest_perimeter(length, width, height) + volume(length, width, height)
	}

	fmt.Printf("%d\n%d\n", wrapping_paper, ribbon)
}