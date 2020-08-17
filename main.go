package main

import (
	"flag"
	"github.com/lingt-xyz/goTool/csv"
	"github.com/lingt-xyz/goTool/csv/fn2fn"
	shrink2 "github.com/lingt-xyz/goTool/csv/shrink"
	matrix2 "github.com/lingt-xyz/goTool/matrix"
)

var (
	normalize   bool
	voc         bool
	index       int
	duplicate   bool
	removeEmpty bool
	indices     string
	filter      bool
	min         int
	max         int
	merge       bool
	f1          string
	f2          string
	shrink      bool
	percentage  float64
	hasHeader   bool
	separate    bool
	rotate      bool
	matrix      bool
	matrixLSTM  bool
	input       string
)

func init() {
	// normalize ARM
	flag.BoolVar(&normalize, "normalize", false, "normalize ARM instructions")

	// generate vocabulary
	flag.BoolVar(&voc, "voc", false, "generate vocabulary from a given file")
	flag.IntVar(&index, "index", -1, "column in the given file to generate vocabulary from")

	// remove duplicate words
	flag.BoolVar(&duplicate, "duplicate", false, "remove duplicates from vocabulary file")

	// remove rows containing empty fields
	flag.BoolVar(&removeEmpty, "removeEmpty", false, "remove rows containing empty fields from a CSV file")
	flag.StringVar(&indices, "indices", "", "Indices of fields need to be checked on emptiness; if this is not provided, all fields are going to be checked.")

	// remove rows that contains less or more tokens that the limitation
	flag.BoolVar(&filter, "filter", false, "remove rows containing unexpected number of tokens")
	flag.IntVar(&min, "min", 1, "minimum number of tokens")
	flag.IntVar(&max, "max", -1, "maximum number of tokens")

	// map ARM and X86 by merging them
	flag.BoolVar(&merge, "merge", false, "remove rows containing empty fields from a CSV file")
	flag.StringVar(&f1, "f1", "ARM.csv", "first file to be merged with the second file")
	flag.StringVar(&f2, "f2", "X86.csv", "second file to be merged with the first file")

	// shrink file
	flag.BoolVar(&shrink, "shrink", false, "shrink file")
	flag.Float64Var(&percentage, "p", 0.5, "percentage of the file to keep")
	flag.BoolVar(&hasHeader, "hasHeader", false, "keep headers")
	flag.BoolVar(&separate, "s", false, "keep both files")

	// rotate dataset folder
	flag.BoolVar(&rotate, "rotate", false, "rotate dataset folder")

	// log matrix
	flag.BoolVar(&matrix, "matrix", false, "log matrix")
	flag.BoolVar(&matrixLSTM, "matrixLSTM", false, "log matrix")
	flag.StringVar(&input, "input", "", "input file name")
	flag.Parse()
}

func main() {
	if normalize {
		normalizeARM(input)
	} else if voc {
		generateVocabulary(input, index)
	} else if duplicate {
		removeDuplicates(input)
	} else if removeEmpty {
		file, indices := csv.GetParameters(input, indices)
		csv.RemoveEmpty(file, indices, hasHeader)
	} else if filter {
		file, indices := csv.GetParameters(input, indices)
		csv.FilterCSV(file, indices, min, max, hasHeader)
	} else if merge {
		fn2fn.MapFunctionsX86AndArm(f1, f2)
	} else if shrink {
		shrink2.ShrinkFile(input, percentage, hasHeader, separate)
	} else if rotate {
		rotateFolder(input)
	} else if matrix {
		matrix2.Matrix(input)
	} else if matrixLSTM {
		matrix2.MatrixLSTM(input)
	}

}
