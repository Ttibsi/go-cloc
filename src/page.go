package src

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type page struct {
	lang         string
	lines        int
	commentLines int
	blanks       int
}

func buildPage(file string) page {
	lang := checkLang(file)
	nums, err := getLength(file, lang)
	if err != nil {
		fmt.Println(err.Error())
	}

	return page{lang, nums[0], nums[1], nums[2]}
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

func getLength(file string, lang string) ([]int, error) {
	if _, ok := LANGUAGES_COMMENTS[lang]; !ok {
		return []int{}, errors.New("Language not in list: " + lang)
	}

	f, err := os.Open(file)
	if err != nil {
		return []int{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineCount := 0
	commentCount := 0
	blankCount := 0
	isMultiline := false

	for scanner.Scan() {
		lineCount++

		if scanner.Text() == "" {
			// blank
			blankCount += 1
			continue
		} else if strings.HasPrefix(scanner.Text(), LANGUAGES_COMMENTS[lang][0]) {
			// Line comment

			// Some file types don't have proper inline comments (md)
			if LANGUAGES_COMMENTS[lang][0] == "" {
				continue
			}

			commentCount += 1
			continue
		} else if len(LANGUAGES_COMMENTS[lang]) == 1 {
			// No multiline comments in this language
			continue
		}

		// These two are separate if statements becasue you can have a multiline comment
		// open and close on the same line
		if strings.HasPrefix(scanner.Text(), LANGUAGES_COMMENTS[lang][1]) {
			isMultiline = true
		}

		if strings.HasSuffix(scanner.Text(), LANGUAGES_COMMENTS[lang][2]) {
			isMultiline = false
			commentCount += 1
		}

		if isMultiline {
			commentCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return []int{}, err
	}

	return []int{lineCount, commentCount, blankCount}, nil
}
