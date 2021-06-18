package main

import (
	"flag"
	"fmt"
	"os"

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
	var info bool
	flag.BoolVar(&info, "info", false, "print info")
	flag.Parse()
	if info {
		fmt.Println("vBrainFuck - BF to C compiler")
		fmt.Println("Build Info: " + CommitHash + " @ " + BuildTime)
		os.Exit(0)
	}

	bf.CompileFile(iFile)
}
