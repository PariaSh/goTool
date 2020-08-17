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

type rowLSTM struct {
	epoch        string
	totalTime    string
	timePerStep  string
	loss         string
	tp           string
	fp           string
	tn           string
	fn           string
	accuracy     string
	precision    string
	recall       string
	auc          string
	valLost      string
	valTp        string
	valFp        string
	valTn        string
	valFn        string
	valAccuracy  string
	valPrecision string
	valRecall    string
	valAuc       string
}

// toColumns converts a row to a list of string to be written to a CSV file
func (r *rowLSTM) toColumns() []string {
	return []string{r.epoch, r.totalTime, r.timePerStep, r.loss, r.tp, r.fp, r.tn, r.fn, r.accuracy, r.precision, r.recall, r.auc, r.valLost, r.valTp, r.valFp, r.valTn, r.valFn, r.valAccuracy, r.valPrecision, r.valRecall, r.valAuc}
}

func MatrixLSTM(input string) {
	if input == "" {
		log.Fatalf("No input log file was given.")
	}
	rows := readLogLSTM(input)
	writeCSVLSTM(rows, input+".csv")
}

// readLog reads a log file into a list of `row`
func readLogLSTM(filePath string) []rowLSTM {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([]rowLSTM, 0, 2<<5)
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

		//251/251 [==============================]
		//1521s 6s/step
		//loss: 1.7734
		//tp: 3902.0000
		//fp: 1031.0000
		//tn: 23292260352.0000
		//fn: 439577.0000
		//accuracy: 1.0000
		//precision: 0.7910
		//recall: 0.0088
		//auc: 0.5297
		//val_loss: 1.5411
		//val_tp: 0.0000e+00
		//val_fp: 120.0000
		//val_tn: 2908262144.0000
		//val_fn: 50809.0000
		//val_accuracy: 1.0000
		//val_precision: 0.0000e+00
		//val_recall: 0.0000e+00
		//val_auc: 0.5142
		ss := strings.Split(scanner.Text(), "-")
		timeS := strings.Split(strings.TrimSpace(ss[1]), " ")
		totalTimeS := strings.TrimSpace(timeS[0])
		totalTime := totalTimeS[:len(totalTimeS)-1]
		timePerStepS := strings.TrimSpace(timeS[1])
		timePerStep := timePerStepS[:len(timePerStepS)-len("s/step")]
		lossS := strings.TrimSpace(ss[2])
		loss := lossS[len("loss: "):]
		tpS := strings.TrimSpace(ss[3])
		tp := tpS[len("tp: "):]
		fpS := strings.TrimSpace(ss[4])
		fp := fpS[len("fp: "):]
		tnS := strings.TrimSpace(ss[5])
		tn := tnS[len("tn: "):]
		fnS := strings.TrimSpace(ss[6])
		fn := fnS[len("fn: "):]
		accuracyS := strings.TrimSpace(ss[7])
		accuracy := accuracyS[len("accuracy: "):]
		precisionS := strings.TrimSpace(ss[8])
		precision := precisionS[len("precision: "):]
		recallS := strings.TrimSpace(ss[9])
		recall := recallS[len("recall: "):]
		aucS := strings.TrimSpace(ss[10])
		auc := aucS[len("auc: "):]
		valLostS := strings.TrimSpace(ss[11])
		valLost := valLostS[len("val_loss: "):]
		valtpS := strings.TrimSpace(ss[12])
		valtp := valtpS[len("val_tp: "):]
		valfpS := strings.TrimSpace(ss[13])
		valfp := valfpS[len("val_fp: "):]
		valtnS := strings.TrimSpace(ss[14])
		valtn := valtnS[len("val_tn: "):]
		valfnS := strings.TrimSpace(ss[15])
		valfn := valfnS[len("val_fn: "):]
		valAccuracyS := strings.TrimSpace(ss[16])
		valAccuracy := valAccuracyS[len("val_accuracy: "):]
		valprecisionS := strings.TrimSpace(ss[17])
		valprecision := valprecisionS[len("val_precision: "):]
		valrecallS := strings.TrimSpace(ss[18])
		valrecall := valrecallS[len("val_recall: "):]
		valaucS := strings.TrimSpace(ss[19])
		valauc := valaucS[len("val_auc: "):]
		rows = append(rows, rowLSTM{
			epoch:        epoch,
			totalTime:    totalTime,
			timePerStep:  timePerStep,
			loss:         loss,
			tp:           tp,
			fp:           fp,
			tn:           tn,
			fn:           fn,
			accuracy:     accuracy,
			precision:    precision,
			recall:       recall,
			auc:          auc,
			valLost:      valLost,
			valTp:        valtp,
			valFp:        valfp,
			valTn:        valtn,
			valFn:        valfn,
			valAccuracy:  valAccuracy,
			valPrecision: valprecision,
			valRecall:    valrecall,
			valAuc:       valauc,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rows
}

// writeCSV writes a header and `rows` into the file `filePath`
func writeCSVLSTM(rows []rowLSTM, filePath string) {
	//headers := row{"epoch", "total time", "time per step", "loss", "accuracy", "val_lost", "val_accuracy"}
	output, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Cannot create output file, got error: %v", err)
	}
	defer output.Close()
	writer := csv.NewWriter(output)
	defer writer.Flush()
	//writer.Write(headers.toColumns())
	for i := range rows {
		writer.Write(rows[i].toColumns())
	}
}
