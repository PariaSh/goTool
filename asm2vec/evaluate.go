package asm2vec

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

const ComparisonSeparator = "_"
const ResultSeparator = ": "

type EvaluateResult struct {
	TP        int
	FP        int
	FN        int
	precision float64
	recall    float64
}

func EvaluateAverage(input string) {
	resultMap := Evaluate(input)
	for _, v := range []string{"bcf", "fla", "sub"} {
		if v1, ok := resultMap["original_"+v]; ok {
			v2 := resultMap[v+"_original"]
			log.Printf("Original v.s. %v: precision:%.4f, recall:%.4f", v,(v1.precision+v2.precision)/2, (v1.recall+v2.recall)/2)
		}
	}

}

func Evaluate(input string) map[string]EvaluateResult {
	fns, logs := getFiles(input)
	fnMap := readFunction(input, fns)
	logMap := readLog(input, logs)

	resultMap := make(map[string]EvaluateResult)
	for k, v := range logMap {
		ss := strings.Split(k, ComparisonSeparator)
		TP := v.matched // this is a match, and it's true
		FP := v.wrongMatch // there is a match, but it's false
		FN := 0
		for _, targetName := range v.noMatchFns { // there is no match
			fns := fnMap[ss[0]]
			for _, sourceFn := range fns {
				if targetName == sourceFn.name {// but it's false
					if sourceFn.noBasicBlock >= 5 {
						FN += 1
					}
				}
			}
		}
		precision := float64(TP) / float64(TP+FP)
		recall := float64(TP) / float64(TP+FN)
		resultMap[k] = EvaluateResult{
			TP:        TP,
			FP:        FP,
			FN:        FN,
			precision: precision,
			recall:    recall,
		}
		log.Printf("%v, TP:%v, FP:%v, FN:%v, precision:%.4f, recall:%.4f", k, TP, FP, FN, precision, recall)
	}
	return resultMap
}

type FnResult struct {
	noMatch    int
	matched    int
	wrongMatch int
	noMatchFns []string
}

func readLog(input string, logs [][]string) map[string]FnResult {
	fnMap := make(map[string]FnResult, len(logs))
	for _, fs := range logs {
		f := fs[0] + ComparisonSeparator + fs[1]
		lines := readLines(path.Join(input, f))
		noMatch, matched, wrongMatch, noMatchFns := splitRecords(lines)
		fnMap[f] = FnResult{
			noMatch:    noMatch,
			matched:    matched,
			wrongMatch: wrongMatch,
			noMatchFns: noMatchFns,
		}
	}
	return fnMap
}

func splitRecords(lines []string) (int, int, int, []string) {
	noMatch, _ := strconv.Atoi(strings.Split(lines[0], ResultSeparator)[1])
	noMatchFns := make([]string, 0)
	matched, _ := strconv.Atoi(strings.Split(lines[1], ResultSeparator)[1])
	wrongMatch, _ := strconv.Atoi(strings.Split(lines[3], ResultSeparator)[1])
	flag := false
	for _, l := range lines {
		if strings.Contains(l, "did not find a match------------------------------------------------------------------") {
			flag = true
			continue
		}
		if strings.Contains(l, "matched expected------------------------------------------------------------------") {
			break
		}
		if flag {
			noMatchFns = append(noMatchFns, l)
		}
	}
	return noMatch, matched, wrongMatch, noMatchFns
}

type Fn struct {
	name         string
	noBasicBlock int
}

func readFunction(input string, fns []string) map[string][]Fn {
	fnMap := make(map[string][]Fn, len(fns))
	for _, f := range fns {
		lines := readLines(path.Join(input, f))
		for _, l := range lines {
			if l == "" {
				continue
			}
			ss := strings.Split(l, " ")
			no, _ := strconv.Atoi(ss[1])
			fnMap[f] = append(fnMap[f], Fn{
				name:         ss[0],
				noBasicBlock: no,
			})
		}
	}
	return fnMap
}

func getFiles(input string) ([]string, [][]string) {
	files, err := ioutil.ReadDir(input)
	if err != nil {
		log.Fatal(err)
	}
	fns := make([]string, 0)
	logs := make([][]string, 0)
	for _, f := range files {
		ss := strings.Split(f.Name(), ComparisonSeparator)
		if len(ss) == 2 {
			logs = append(logs, ss)
		} else {
			fns = append(fns, f.Name())
		}
	}
	return fns, logs
}

func readLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file %v", fileName)
	}
	defer file.Close()

	lines := make([]string, 0)
	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err == nil {
	//		lines = append(lines, line)
	//	} else {
	//		break
	//	}
	//}
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}
