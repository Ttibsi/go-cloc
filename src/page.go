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
	whitespace   int // Maybe
}

func buildPage(file string) page {
	lang := checkLang(file)
	lines := getLength(file)

	return page{lang: lang, lines: lines}
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

func getLength(file string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineCount := 0

	//TODO: Add a way here to scan for commented out lines - need to work out how
	// to detect the comment character here
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return -1
	}

	return lineCount
}
