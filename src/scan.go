package src

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func getAllFilePaths(
	startingDir string,
	excludeDir []string,
	excludeRegex []string,
) ([]string, error) {
	ret := []string{}

	criteria := func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for _, regex := range excludeRegex {
			match, _ := regexp.MatchString(regex, path)
			if match {
				return nil
			}
		}

		for _, dir := range excludeDir {
			if strings.Contains(path, dir) {
				return nil
			}
		}

		ret = append(ret, path)

		return nil
	}

	err := filepath.WalkDir(startingDir, criteria)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func printFormatter(pages []page) string { return "" }

func StartLogic(f Flags) {
	if f.Version == true {
		fmt.Println("Version: ")
		os.Exit(0)
	}
	// TODO: git-dir
	// TODO: No-recurse
	// TODO: find-lang
	var files []string
	for _, dir := range f.Directory {
		ret, err := getAllFilePaths(dir, f.ExcludeDir, f.Exclude)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		files = append(files, ret...)
	}

	var pages []page
	for _, file := range files {
		pages = append(pages, buildPage(file))
	}

	// Read all data and create the output string, and print it
	fmt.Println(printFormatter(pages))
}
