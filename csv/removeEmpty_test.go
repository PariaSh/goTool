package csv

import "testing"

func TestPrintFiledIndices(t *testing.T) {
	header := []string{"libraryName", "binaryName", "functionName",
		"version", "architecture", "compiler", "optimization", "obfuscation", "edgeCoverage",
		"version", "architecture", "compiler", "optimization", "obfuscation", "edgeCoverage",
	}
	indices := []int{0,1,2,5,100}
	printFiledIndices(header, indices)
}
