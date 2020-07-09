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
	duplicate   bool
	removeEmpty bool
	indices     string
	merge       bool
	f1          string
	f2          string
	shrink      bool
	percentage  float64
	keepHeader  bool
	separate    bool
	rotate      bool
	matrix      bool
	input       string
)

func init() {
	// normalize ARM
	flag.BoolVar(&normalize, "normalize", false, "normalize ARM instructions")

	// generate vocabulary
	flag.BoolVar(&voc, "voc", false, "generate vocabulary")

	// remove duplicate words
	flag.BoolVar(&duplicate, "duplicate", false, "remove duplicates from vocabulary file")

	// remove rows containing empty fields
	flag.BoolVar(&removeEmpty, "removeEmpty", false, "remove rows containing empty fields from a CSV file")
	flag.StringVar(&indices, "indices", "", "Indices of fields need to be checked on emptiness; if this is not provided, all fields are going to be checked.")

	// map ARM and X86 by merging them
	flag.BoolVar(&merge, "removeEmpty", false, "remove rows containing empty fields from a CSV file")
	flag.StringVar(&f1, "f1", "ARM.csv", "first file to be merged with the second file")
	flag.StringVar(&f2, "f2", "X86.csv", "second file to be merged with the first file")

	// shrink file
	flag.BoolVar(&shrink, "shrink", false, "shrink file")
	flag.Float64Var(&percentage, "p", 0.5, "percentage of the file to keep")
	flag.BoolVar(&keepHeader, "h", false, "keep headers")
	flag.BoolVar(&separate, "s", false, "keep both files")

	// rotate dataset folder
	flag.BoolVar(&rotate, "rotate", false, "rotate dataset folder")

	// log matrix
	flag.BoolVar(&matrix, "matrix", false, "log matrix")
	flag.StringVar(&input, "input", "", "input file name")
	flag.Parse()
}

func main() {
	if normalize {
		normalizeARM(input)
	} else if voc {
		generateVocabulary(input)
	} else if duplicate {
		removeDuplicates(input)
	} else if removeEmpty {
		file, indices := csv.GetParameters(input, indices)
		csv.UpdateCSV(file, indices)
	} else if merge {
		fn2fn.MapFunctionsX86AndArm(f1, f2)
	} else if shrink {
		shrink2.ShrinkFile(input, percentage, keepHeader, separate)
	} else if rotate {
		rotateFolder(input)
	} else if matrix {
		matrix2.Matrix(input)
	}

}
