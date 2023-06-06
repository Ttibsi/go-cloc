package src

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type page struct {
	lang         string
	lines        int
	commentLines int
}

func buildPage(file string) page {
	lang := checkLang(file)
	lines, comments := getLength(file, lang)

	return page{lang: lang, lines: lines, commentLines: comments}
}

func checkLang(file string) string {
	if filepath.Ext(file) != "" {
		return filepath.Ext(file)
	}

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line == 1 {
			if scanner.Text()[0:2] == "#!/" {
				return string(strings.Fields(scanner.Text())[len(scanner.Text())-1])
			}
		}
	}

	return ""
}

func getLength(file string, lang string) (int, int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return -1, -1
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineCount := 0
	commentCount := 0
	isMultiline := false

	//TODO: Add a way here to scan for commented out lines - need to work out how
	// to detect the comment character here
	for scanner.Scan() {
		// Line comment
		if strings.HasPrefix(scanner.Text(), LANGUAGES_COMMENTS[lang][0]) {
			commentCount++
		}

		lineCount++

		// No multiline comments in this language
		if len(LANGUAGES_COMMENTS[lang]) == 1 {
			break
		}

		if strings.HasPrefix(scanner.Text(), LANGUAGES_COMMENTS[lang][1]) {
			// commentCount += checkMultiLine(scanner.Text(), lang)
			isMultiline = true
		} else if strings.HasPrefix(scanner.Text(), LANGUAGES_COMMENTS[lang][2]) {
			isMultiline = false
		}

		if isMultiline {
			commentCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return -1, -1
	}

	return lineCount, commentCount
}
