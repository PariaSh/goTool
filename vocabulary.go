package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

func generateVocabulary(fileName string) {
	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	vocabularyMap := make(map[string]int, 1<<10)

	reader := csv.NewReader(sourceFile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		edgeCoverage := record[8]
		bb := strings.Split(edgeCoverage, "#")
		for i := range bb {
			token := strings.TrimSpace(bb[i])
			if token != "" {
				vocabularyMap[token]++
			}
		}
	}

	targetFileName := fileName + ".voc"
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("Failed to create target file: %v, got error: %v", targetFileName, err)
	}
	defer targetFile.Close()

	for k, _ := range vocabularyMap {
		targetFile.WriteString(k + "\n")
	}
}

func removeDuplicates(fileName string) {
	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	targetFileName := fileName + ".duplicatesRemoved"
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("Failed to create target file: %v, got error: %v", targetFileName, err)
	}
	defer targetFile.Close()

	reader := bufio.NewReader(sourceFile)

	set := make(map[string]bool) // New empty set
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		if s == ""{
			continue
		}
		set[s] = true // Add
	}

	for k := range set { // Loop
		_, _ = targetFile.WriteString(k)
	}
}
