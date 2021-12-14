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
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parsed the provided csv file.")
	}
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
