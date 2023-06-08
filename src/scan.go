package src

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

		if !d.IsDir() {
			ret = append(ret, path)
		}

		return nil
	}

	err := filepath.WalkDir(startingDir, criteria)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func printFormatter(pages []page) string {
	// {"langname": []int{files, lines, comments, blanks}}
	out := make(map[string][]int)
	var totalLines int
	var totalComment int
	var totalBlank int
	var totalCode int

	for _, page := range pages {
		if exist, ok := out[page.lang]; ok {
			currFileCount := exist[0]
			for i, newVal := range []int{currFileCount + 1, page.lines, page.commentLines, page.blanks} {
				exist[i] += newVal
			}
		} else {
			if slices.Contains(BLACKLIST, page.lang) {
				continue
			}
			out[page.lang] = []int{1, page.lines, page.commentLines, page.blanks}
		}
	}

	var builder strings.Builder
	builder.WriteString("\t" + strconv.Itoa(len(pages)) + " files scanned\n")
	// TODO: Print time of exexcution
	builder.WriteString("\n") //empty line in output
	builder.WriteString("| Language | Files | Lines | Comments | Blank lines | Code lines |")
	builder.WriteString("\n") //empty line in output
	builder.WriteString("+----------+-------+-------+----------+-------------+------------+")
	builder.WriteString("\n") //empty line in output
	// TODO: Rework getting file names/types to print

	for k, v := range out {
		wordlen := len(k)
		codeLines := v[1] - (v[2] + v[3])
		langName := cases.Title(language.Und, cases.NoLower).String(k)

		totalLines += v[1]
		totalComment += v[2]
		totalBlank += v[3]
		totalCode += codeLines

		builder.WriteString("| " + langName + strings.Repeat(" ", len("language")-wordlen) + " ")
		builder.WriteString(
			"| " + strings.Repeat(
				" ",
				(len("files")-len(strconv.Itoa(v[0]))),
			) + strconv.Itoa(
				v[0],
			) + " ",
		)
		builder.WriteString(
			"| " + strings.Repeat(
				" ",
				(len("lines")-len(strconv.Itoa(v[1]))),
			) + strconv.Itoa(
				v[1],
			) + " ",
		)
		builder.WriteString(
			"| " + strings.Repeat(
				" ",
				(len("comments")-len(strconv.Itoa(v[2]))),
			) + strconv.Itoa(
				v[2],
			) + " ",
		)
		builder.WriteString(
			"| " + strings.Repeat(
				" ",
				(len("blank lines")-len(strconv.Itoa(v[3]))),
			) + strconv.Itoa(
				v[3],
			) + " ",
		)
		builder.WriteString(
			"| " + strings.Repeat(
				" ",
				(len("code lines")-len(strconv.Itoa(codeLines))),
			) + strconv.Itoa(
				codeLines,
			) + " ",
		)
		builder.WriteString("|\n") //empty line in output
	}

	builder.WriteString("+----------+-------+-------+----------+-------------+------------+")
	builder.WriteString("\n") //empty line in output

	builder.WriteString(
		"|   Total  | " +
			strings.Repeat(" ", len("files")-len(strconv.Itoa(len(pages)))) +
			strconv.Itoa(len(pages)) +
			" | " +
			strings.Repeat(" ", len("lines")-len(strconv.Itoa(totalLines))) +
			strconv.Itoa(totalLines) +
			" | " +
			strings.Repeat(" ", len("comments")-len(strconv.Itoa(totalComment))) +
			strconv.Itoa(totalComment) +
			" | " +
			strings.Repeat(" ", len("blank lines")-len(strconv.Itoa(totalBlank))) +
			strconv.Itoa(totalBlank) +
			" | " +
			strings.Repeat(" ", len("code lines")-len(strconv.Itoa(totalCode))) +
			strconv.Itoa(totalCode) +
			" | ")
	builder.WriteString("\n") //empty line in output
	builder.WriteString("+----------+-------+-------+----------+-------------+------------+")

	return builder.String()
}

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
