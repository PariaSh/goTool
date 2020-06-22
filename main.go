package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
)

func main() {
	f := flag.String("f", "", "input file name")
	flag.Parse()
	normalized(*f)
	//removeDuplicates(*f)
}

var reMem = regexp.MustCompile(`MEM\w+`)
var reTmp = regexp.MustCompile(`t\d+`)
func normalized(fileName string) {
	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	targetFileName := fileName + ".normalized"
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("Failed to create target file: %v, got error: %v", targetFileName, err)
	}
	defer targetFile.Close()

	reader := bufio.NewReader(sourceFile)

	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		s = reMem.ReplaceAllString(s, "MEM")
		s = reTmp.ReplaceAllString(s, "TMP")
		_, _ = targetFile.WriteString(s)
	}
}

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
