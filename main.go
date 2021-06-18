package main

import (
	"flag"

	"github.com/vopi181/vBrainFuck/bf"
)

var (
	BuildTime  string
	CommitHash string
	GoVersion  string
)

func main() {
	var iFile string
	flag.StringVar(&iFile, "file", "test.bf", "input file")
	flag.Parse()
	//fmt.Println("vBrainFuck - BF to C compiler")
	//fmt.Println("Build Info: " + BuildTime + "!")
	//fmt.Println("working on " + iFile)
	bf.CompileFile(iFile)
}
