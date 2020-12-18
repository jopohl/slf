package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/jopohl/slf/src"
)

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage %s PATH\nArguments:\n", os.Args[0])
		flag.PrintDefaults()
	}

	numFiles := flag.Int("n", 10, "Number of entries to display. Must be a positive number.")
	dirMode := flag.Bool("d", false, "Scan for directories instead of files.")
	numThreads := flag.Int("x", runtime.NumCPU()+1, "Number of threads to use.")
	flag.Parse()

	if *numFiles <= 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pathsToScan := flag.Args()
	if len(pathsToScan) < 1 {
		pathsToScan = []string{"."}
	}

	for _, path := range pathsToScan {
		tree := slf.ScanDirectoryTree(path, *numThreads)
		entries := make([]*slf.Node, *numFiles)
		slf.UpdateSizes(tree, &entries, *dirMode)
		slf.PrintSizes(entries, *numFiles)
	}
}
