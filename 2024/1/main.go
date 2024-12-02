package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// open input file
	readFile, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var left []int
	var right []int

	for _, line := range fileLines {
		//fmt.Printf("line is %s\n", line)

		split := strings.Split(line, "   ")
		numLeft, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Error converting string:", err)
		}
		numRight, err := strconv.Atoi(split[1])
		if err != nil {
			// Handle the error, e.g., print it
			fmt.Println("Error converting string:", err)
		}
		left = append(left, numLeft)
		right = append(right, numRight)

	}

	sort.Ints(left)
	sort.Ints(right)

	var answer float64

	for i, l := range left {

		a := float64(l - right[i])
		fmt.Printf("Left %d Right %d, Diff %f\n", l, right[i], math.Abs(a))
		answer = answer + math.Abs(a)
	}

	fmt.Printf("Answer is %f\n", answer)
}
