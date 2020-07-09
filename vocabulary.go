package main

import (
	"bufio"
	"log"
	"os"
)

func removeDuplicates(fileName string){
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

		set[s] = true            // Add
	}

	for k := range set {         // Loop
		_, _ = targetFile.WriteString(k)
	}
}
