package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetParameters(input, indexes string) (string, []int) {
	if input == "" {
		log.Fatal("No input CSV file was given.")
	}
	indices := make([]int, 0, 2)
	if indexes == "" {
		log.Printf("indexes are not specified, all fields are going to be checked.")
	} else {
		for _, i := range strings.Split(indexes, ",") {
			index, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				log.Fatalf("Indexes should be of the format: \"number[, number]\"; but got: %v", indexes)
			}
			indices = append(indices, index)
		}
	}
	return input, indices
}

// UpdateCSV removes rows that have empty fields at `indices`
func RemoveEmpty(filePath string, indices []int, hasHeader bool) {
	inputCSV, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Cannot open file %v, got error: %v", filePath, err)
	}
	defer inputCSV.Close() // this needs to be after the err check

	// write the result to two file so it's easy to validate empty fields.
	outputNonEmptyCSV, err1 := os.Create(filePath + "-non-empty")
	outputEmptyCSV, err2 := os.Create(filePath + "-empty")
	if err1 != nil || err2 != nil {
		log.Fatalf("Cannot create output file, got error: %v", err)
	}
	defer outputNonEmptyCSV.Close()
	defer outputEmptyCSV.Close()

	nonEmptyWriter := csv.NewWriter(outputNonEmptyCSV)
	emptyWriter := csv.NewWriter(outputEmptyCSV)
	defer nonEmptyWriter.Flush()
	defer emptyWriter.Flush()

	reader := csv.NewReader(inputCSV)
	if hasHeader {
		headers, _ := reader.Read()
		nonEmptyWriter.Write(headers)
		emptyWriter.Write(headers)

		printFiledIndices(headers, indices)
	}
row:
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(indices) == 0 {
			for i := range record {
				if strings.TrimSpace(record[i]) == "" {
					_ = emptyWriter.Write(record)
					continue row
				}
			}
			_ = nonEmptyWriter.Write(record)
		} else {
			for i := range record {
				// doing it in this way can avoid index of range
				for _, j := range indices {
					if i == j {
						if strings.TrimSpace(record[i]) == "" {
							_ = emptyWriter.Write(record)
							continue row
						}
					}
				}
			}
			_ = nonEmptyWriter.Write(record)
		}
	}
}

// printFiledIndices prints which fields has user chosen to check
func printFiledIndices(header []string, indices []int) {
	log.Printf("You checked the following fileds:")
	fmt.Printf("-------------------\n")
	for i, e := range header {
		fmt.Printf("|%-15s|", e)
		if len(indices) == 0 {
			fmt.Print("Y")
		} else {
			for _, j := range indices {
				if i == j {
					fmt.Print("Y")
				}
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("-------------------\n")
}

// UpdateCSV removes rows that have empty fields at `indices`
func FilterCSV(filePath string, indices []int, min, max int, hasHeader bool) {
	inputCSV, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Cannot open file %v, got error: %v", filePath, err)
	}
	defer inputCSV.Close() // this needs to be after the err check

	// write the result to two file so it's easy to validate the correctness.
	qualifiedFile, err1 := os.Create(filePath + "-qualified")
	nonQualifiedFile, err2 := os.Create(filePath + "-non-qualified")
	if err1 != nil || err2 != nil {
		log.Fatalf("Cannot create output file, got error: %v", err)
	}
	defer qualifiedFile.Close()
	defer nonQualifiedFile.Close()

	qualifiedWriter := csv.NewWriter(qualifiedFile)
	nonQualifiedWriter := csv.NewWriter(nonQualifiedFile)
	defer qualifiedWriter.Flush()
	defer nonQualifiedWriter.Flush()

	reader := csv.NewReader(inputCSV)
	if hasHeader {
		headers, _ := reader.Read()
		qualifiedWriter.Write(headers)
		nonQualifiedWriter.Write(headers)
		printFiledIndices(headers, indices)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		tokens := make([]string, 0)
		for i := range record {
			// doing it in this way can avoid index of range
			for _, j := range indices {
				if i == j {
					tokens = append(tokens, strings.Split(record[i], "#")...)
				}
			}
		}
		if min <= len(tokens) {
			if max == -1 {
				_ = qualifiedWriter.Write(record)
			} else if len(tokens) <= max {
				_ = qualifiedWriter.Write(record)
			} else {
				_ = nonQualifiedWriter.Write(record)
			}
		} else {
			_ = nonQualifiedWriter.Write(record)
		}
	}
}
