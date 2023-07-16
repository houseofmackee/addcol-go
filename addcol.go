package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kingpin/v2"
)

var (
	argColumn     = kingpin.Arg("column", "Column number where to insert Value (9999 to add column at the end)").Required().Int()
	argValue      = kingpin.Arg("value", "Value to insert").Required().String()
	argInputFile  = kingpin.Arg("infile", "Input CSV file name").Required().ExistingFile()
	argOutputFile = kingpin.Arg("outfile", "Output CSV file name (defaults to stdout)").String()
)

func main() {
	// parse command line arguments
	kingpin.Parse()

	// do we have a valid column number?
	column := *argColumn
	if column < 1 {
		fmt.Println("Invalid column number")
		os.Exit(1)
	}

	// open input fileName
	fileName := *argInputFile
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", fileName)
		os.Exit(1)
	}
	defer file.Close()

	// read input file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading file %s\n", os.Args[3])
		log.Fatal(err)
	}

	// check if file is empty
	if len(records) == 0 {
		fmt.Printf("File %s is empty\n", fileName)
		os.Exit(1)
	}

	// special case: add column at the end
	numCols := len(records[0])
	if column == 9999 {
		column = numCols + 1
	}

	// check if all lines have enough columns and if all lines have the
	// same number of columns (based on first line)
	for i, record := range records {
		if len(record) < column-1 {
			fmt.Printf("Line %d: not enough columns\n", i+1)
			os.Exit(1)
		}
		if len(record) != numCols {
			fmt.Printf("Line %d: number of columns differs from first line\n", i+1)
			os.Exit(1)
		}
	}

	// add column at position col with value argValue
	insertValue := *argValue
	for i, record := range records {
		//records[i] = append(record[:col-1], append([]string{*argValue}, record[col-1:]...)...)
		records[i] = append(record, "")
		copy(records[i][column:], records[i][column-1:])
		records[i][column-1] = insertValue
	}

	// open output file or use stdout as default
	outFile := os.Stdout
	if *argOutputFile != "" {
		outFile, err = os.OpenFile(*argOutputFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening output file %s: %s\n", *argOutputFile, err)
		}
		defer outFile.Close()
	}

	// write output file
	writer := csv.NewWriter(outFile)
	writer.WriteAll(records)
	if err := writer.Error(); err != nil {
		log.Fatalln("Error writing csv:", err)
	}
}
