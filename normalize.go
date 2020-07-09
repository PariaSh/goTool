package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

var reMem = regexp.MustCompile(`MEM\w+`)
var reTmp = regexp.MustCompile(`t\d+`)

func normalizeARM(fileName string) {
	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	targetFileName := fileName + ".normalizeARM"
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
