package matrix

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type row struct {
	epoch       string
	totalTime   string
	timePerStep string
	loss        string
	accuracy    string
	valLost     string
	valAccuracy string
}

// toColumns converts a row to a list of string to be written to a CSV file
func (r *row) toColumns() []string {
	return []string{r.epoch, r.totalTime, r.timePerStep, r.loss, r.accuracy, r.valLost, r.valAccuracy}
}

func Matrix(input string) {
	if input == "" {
		log.Fatalf("No input log file was given.")
	}
	rows := readLog(input)
	writeCSV(rows, input+".csv")
}

// readLog reads a log file into a list of `row`
func readLog(filePath string) []row {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([]row, 0, 2<<5)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" { // ignore the ending empty lines
			continue
		}
		epoch := strings.Split(strings.Split(scanner.Text(), " ")[1], "/")[0]

		// every two lines construct a row, so read two line in every loop
		if !scanner.Scan() {
			log.Printf("Incompelte file %q", filePath)
		}
		if strings.TrimSpace(scanner.Text()) == "" { // ignore the ending empty lines
			continue
		}
		ss := strings.Split(scanner.Text(), "-")
		timeS := strings.Split(strings.TrimSpace(ss[1]), " ")
		totalTimeS := strings.TrimSpace(timeS[0])
		totalTime := totalTimeS[:len(totalTimeS)-1]
		timePerStepS := strings.TrimSpace(timeS[1])
		timePerStep := timePerStepS[:len(timePerStepS)-len("s/step")]
		lossS := strings.TrimSpace(ss[2])
		loss := lossS[len("loss: "):]
		accuracyS := strings.TrimSpace(ss[3])
		accuracy := accuracyS[len("accuracy: "):]
		valLostS := strings.TrimSpace(ss[4])
		valLost := valLostS[len("val_loss: "):]
		valAccuracyS := strings.TrimSpace(ss[5])
		valAccuracy := valAccuracyS[len("val_accuracy: "):]
		rows = append(rows, row{
			epoch:       epoch,
			totalTime:   totalTime,
			timePerStep: timePerStep,
			loss:        loss,
			accuracy:    accuracy,
			valLost:     valLost,
			valAccuracy: valAccuracy,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rows
}

// writeCSV writes a header and `rows` into the file `filePath`
func writeCSV(rows []row, filePath string) {
	headers := row{"epoch", "total time", "time per step", "loss", "accuracy", "val_lost", "val_accuracy"}
	output, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Cannot create output file, got error: %v", err)
	}
	defer output.Close()
	writer := csv.NewWriter(output)
	defer writer.Flush()
	writer.Write(headers.toColumns())
	for i := range rows {
		writer.Write(rows[i].toColumns())
	}
}
