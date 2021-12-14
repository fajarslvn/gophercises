package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// helper flag berguna untuk user yg ingin mengetahui cara pemakaiannya dgn flag '-h'
	// agar bisa digunakan deklarasikan variabel ke 'flag.String("tipe file", "default nama file", "deskripsi")'
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// Add this after all flags are defined and before flags are accessed by the program.
	flag.Parse()

	// open file and handle if there's an error
	file, err := os.Open(*csvFileName)
	if err != nil {
		// If the file error or not found, terminates/close the program immediately
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))
	}
	// read file in format csv
	r := csv.NewReader(file)
	// read all lines and handle if there's an error
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parsed the provided csv file.")
	}
	// print the lines
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	// Conventionally, code zero indicates success, non-zero an error.
	// The program terminates immediately; deferred functions are not run.
	os.Exit(1)
}
