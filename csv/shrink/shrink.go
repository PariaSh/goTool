package shrink

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ShrinkFile(fileName string, percentage float64, hasHeader bool, separate bool){
	if fileName == "" {
		log.Fatalf("No input fileName was given")
	} else {
		log.Printf("Provided file: %v", fileName)
	}

	if percentage < 0 || percentage > 1.0 {
		log.Fatalf("Invalid percentage: expecting (0, 1.0), got %v", percentage)
	}

	shrink(fileName, percentage, hasHeader, separate)
}
func shrink(fileName string, percentage float64, hasHeader, separate bool) {

	sourceFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v, got error: %v", fileName, err)
	}
	defer sourceFile.Close()

	targetFileName := fileName + "." + strconv.FormatFloat(percentage, 'f', 2, 64)
	targetFile, err := os.Create(targetFileName)
	if err != nil {
		log.Fatalf("Failed to create target file: %v, got error: %v", targetFileName, err)
	}
	defer targetFile.Close()

	var separateFile *os.File
	if separate {
		separateFileName := fileName + "." + strconv.FormatFloat(1-percentage, 'f', 2, 64)
		separateFile, err = os.Create(separateFileName)
		if err != nil {
			log.Fatalf("Failed to create target file: %v, got error: %v", separateFile, err)
		}
		defer separateFile.Close()
	}

	// Scanner cannot deal with "too long token"
	//scanner := bufio.NewScanner(sourceFile)
	//var header string
	//var headerScanned bool
	//for scanner.Scan() {
	//	if hasHeader && !headerScanned {
	//		header = scanner.Text()
	//		headerScanned = true
	//		_, _ = targetFile.WriteString(header + "\n")
	//		_, _ = separateFile.WriteString(header + "\n")
	//		continue
	//	}
	//	if randomize(percentage) {
	//		_, _ = targetFile.WriteString(scanner.Text() + "\n")
	//	} else {
	//		if separate {
	//			_, _ = separateFile.WriteString(scanner.Text() + "\n")
	//		}
	//	}
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	log.Fatal(err)
	//}

	reader := bufio.NewReader(sourceFile)

	if hasHeader {
		header, _ := reader.ReadString('\n')
		_, _ = targetFile.WriteString(header)
		_, _ = separateFile.WriteString(header)
	}
	for{
		s, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		if randomize(percentage) {
			_, _ = targetFile.WriteString(s)
		} else {
			if separate {
				_, _ = separateFile.WriteString(s)
			}
		}
	}
}
