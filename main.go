package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func path_trawl(dir string) []string {
	// TODO: flag to enable specific file types currently blacklisted
	var blacklist_file_types = []string{
		".DS_Store",
		".editorconfig",
		".gif",
		".gitattributes",
		".gitignore",
		".gitmodules",
		".jpg",
		".json",
		".md",
		".png",
		".pyc",
		".pyi",
		".rst",
		".txt", //maybe - might want to include cmake files
		".xml",
		".yaml",
	}

	all_paths := make([]string, 1000) // slice of max length 1000

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else {

				for _, elem := range strings.Split(path, "\n") {
					if contains(blacklist_file_types, filepath.Ext(elem)) {
						continue
					} else if strings.Contains(filepath.Dir(elem), ".") {
						continue
					} else if filepath.Base(elem) == "LICENSE" {
						continue
					} else {
						// skip directories, we only want files
						info, err := os.Stat(elem)
						if err != nil {
							log.Panic(err)
						}

						if info.IsDir() {
							continue
						}
					}

					all_paths = append(all_paths, elem)
				}
				return nil
			}
		})
	if err != nil {
		log.Println(err)
	}

	return all_paths
}

func file_length(filepath string) int {
	var length int = 0

	if len(filepath) == 0 {
		return 0
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//TODO: handle commented out lines
		if scanner.Text() != "" { // Ignore empty lines
			length++
		}
	}

	return length
}

func count_through_directory(dir string) (int, int) {
	var count int = 0
	var file_count int = 0

	all_files := path_trawl(dir)

	for _, file := range all_files {
		count += file_length(file)
		if count != 0 {
			fmt.Println(file)
			file_count++
		}
	}

	return count, file_count
}

func output_value(c int, fc int) {
	// TODO: handle output based on file type?
	// Ex: " 12 lines across 2 go files, 16 lines across 6 python files"
	fmt.Println(fmt.Sprintf("Found %d lines of code across %d files", c, fc))
}

func main() {
	var extensions string

	app := &cli.App{
		Name:  "count-loc",
		Usage: "Count lines of code in a given directory",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "enable-ext",
				Value:       "",
				Usage:       "Add file types currently blacklisted",
				Destination: &extensions,
			},
		},
		Action: func(cCtx *cli.Context) error {
			var path string
			if cCtx.NArg() > 0 {
				path = cCtx.Args().Get(0)
				//break here
			} else {
				fmt.Println("Error: No filepath entered")
			}

			fmt.Println(extensions)
			fmt.Println(path)
			//count, file_count := count_through_directory(path)
			//output_value(count, file_count)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
