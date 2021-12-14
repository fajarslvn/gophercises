package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Make a helper flag for user guides in examples '-h' or ''--help'
	// Define flag '-csv' using 'flag.String("helper name", "default file you want to call", "flag description")'
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// Add this after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	// open file and handle if there's an error, in example file not found or miss typing
	file, err := os.Open(*csvFileName)
	if err != nil {
		// If the file error or not found, immediately terminates or close the program
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))
	}
	// read file in csv format
	r := csv.NewReader(file)
	// read all lines and handle if there's an error
	lines, err := r.ReadAll()
	if err != nil {
		// If the file error or not found, immediately terminates or close the program
		exit("Failed to parsed the provided csv file.")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	os.Exit(1)
}
