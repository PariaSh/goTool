package evaluation

import (
	"regexp"
)

//--ImageMagick-7.0.1_10  gmp-6.1.0  openssl-1.0.2s  zlib-1.2.7.1
//SELECT
//LibraryName, Version, BinaryName, Architecture, Compiler, Optimization,
//FunctionName, FunctionName, NoInstructions, NoBasicBlocks, NoEdgeCoverage, NoCallWalk
//FROM Function
//JOIN Binary USING (BinaryID)
//JOIN Library USING (LibraryID)
//WHERE Compiler='gcc'
//AND LibraryName in ('zlib', 'gmp', 'openssl', 'ImageMagick')
//AND Version in ('1.2.7.1', '7.0.1_10','6.1.0', '1.0.2s')

// Number of blocks
// LibraryName, Version, BinaryName, Architecture, Compiler, Optimization, FunctionName, FunctionName, NoInstructions, NoBasicBlocks, NoEdgeCoverage, NoCallWalk

// Library time
// LibraryName, Version, BinaryName, Architecture, Compiler, Optimization, Obfuscation, RunTime

// Function time
//

type vexRow struct {
	libraryName    string
	version        string
	binaryName     string
	architecture   string
	compiler       string
	optimization   string
	obfuscation    string
	functionName   string
	NoInstructions string
	NoBasicBlocks  string
	NoEdgeCoverage string
	NoCallWalk     string
}

func newRecord(records []string) *vexRow {
	return &vexRow{
		libraryName:    records[0],
		version:        records[1],
		binaryName:     records[2],
		architecture:   records[3],
		compiler:       records[4],
		optimization:   records[5],
		functionName:   records[6],
		NoInstructions: records[7],
		NoBasicBlocks:  records[8],
		NoEdgeCoverage: records[9],
		NoCallWalk:     records[10],
	}
}

func newBinary(records []string) *vexRow {
	return &vexRow{
		libraryName:    records[1],
		version:        records[2],
		binaryName:     records[5],
		architecture:   records[3],
		compiler:       "gcc",
		optimization:   records[4],
		obfuscation:    "",
		functionName:   "",
		NoInstructions: "",
		NoBasicBlocks:  "",
		NoCallWalk:     "",
	}
}

var rePrevious = regexp.MustCompile(`_volume_dataset_(\w+)-([0-9\._]+)_(arm|x86)-gcc-(O[0|1|2|3])_(.*so)(.*)`)

func GetBinary(fileName string) {
	newBinary(rePrevious.FindStringSubmatch(fileName))
}
