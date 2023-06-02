package src

import (
	"fmt"
	"os"
)

func getAllFilePaths(startingDir string, excludeDir string, excludeRegex string) []string {
	ret := []string{}

	criteria := func(file string) bool {
		//TODO handle exclusions
		return true
	}

	return ret
}

func printFormatter(pages []page) string { return "" }

func StartLogic(f Flags) {
	if f.Version == true {
		fmt.Println("Version: ")
		os.Exit(0)
	}

	// build a list of all files to scan, ignoring excluded files/directories
	// TODO No-recurse
	// TODO find-lang
	var files []string
	for _, dir := range f.Directory {
		files = append(files, getAllFilePaths(dir, f.ExcludeDir, f.Exclude)...)
	}

	var pages []page
	for _, file := range files {
		// Scan each line, turn into page struct
	}

	// Read all data and create the output string, and print it
	fmt.Println(printFormatter(pages))
}
