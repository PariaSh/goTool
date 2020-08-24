package asm2vec

import "testing"

const Asm2vecTestFolder = "/Users/ling/git/goTool/asm2vec/libgmp"

func TestEvaluateAverage(t *testing.T) {
	EvaluateAverage(Asm2vecTestFolder)
}

func TestEvaluate(t *testing.T) {
	Evaluate(Asm2vecTestFolder)
}

func TestGetFiles(t *testing.T){
	fns, logs := getFiles(Asm2vecTestFolder)
	t.Logf("%+v", fns)
	t.Logf("%+v", logs)
}

func TestReadLog(t *testing.T){
	_, logs := getFiles(Asm2vecTestFolder)
	fnMap := readLog(Asm2vecTestFolder, logs)
	for k, v := range fnMap{
		t.Logf("%+v, %+v", k, v)
	}
}

func TestReadFunction(t *testing.T){
	fns, _ := getFiles(Asm2vecTestFolder)
	fnMap := readFunction(Asm2vecTestFolder, fns)
	for k, v := range fnMap{
		t.Logf("%+v", k)
		for _, f := range v{
			t.Logf("%+v", f)
		}
	}
}
