package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func path_trawl(dir string, setFlags flags) []string {
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
		".swp",
		".tmp",
		".txt", //maybe - might want to include cmake files
		".xml",
		".yaml",
	}

	if len(setFlags.new_blacklist_item) > 0 {
		blacklist_file_types = append(blacklist_file_types, setFlags.new_blacklist_item)
	}

	all_paths := make([]string, 1000) // slice of max length 1000

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else {

				for _, elem := range strings.Split(path, "\n") {
					if !setFlags.enable_all && contains(blacklist_file_types, filepath.Ext(elem)) {
						if setFlags.extension != filepath.Ext(elem) {
							continue
						}
					} else if !setFlags.use_hidden_dirs && strings.Contains(filepath.Dir(elem), ".") {
						continue
					} else if !setFlags.enable_all && filepath.Base(elem) == "LICENSE" {
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

func count_through_directory(dir string, setFlags flags) (int, int) {
	var count int = 0
	var file_count int = 0

	all_files := path_trawl(dir, setFlags)

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
	fmt.Printf("Found %d lines of code across %d files\n", c, fc)
}

type flags struct {
	extension          string
	enable_all         bool
	use_hidden_dirs    bool
	new_blacklist_item string
}

func main() {
	var setFlags flags

	cli.VersionPrinter = func(cCtx *cli.Context) {
		stdout, err := exec.Command("git", "tag").Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		val := strings.Split(string(stdout), "\n")
		fmt.Println(val[0])
		os.Exit(0) // If this is called with other values, make sure it only runs this
	}

	app := &cli.App{
		Name:    "c-loc",
		Usage:   "Count lines of code in a given directory",
		Version: "v0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "enable-ext",
				Value:       "",
				Usage:       "Add file types currently blacklisted",
				Destination: &setFlags.extension,
			},
			&cli.BoolFlag{
				Name:        "enable-all",
				Value:       false,
				Usage:       "Enable searching all extensions blacklisted by default",
				Destination: &setFlags.enable_all,
			},
			&cli.BoolFlag{
				Name:        "use-hidden-dirs",
				Value:       false,
				Usage:       "Include searching through hidden directors, such as .git",
				Destination: &setFlags.enable_all,
			},
			&cli.StringFlag{
				Name:        "ignore-ext",
				Value:       "",
				Usage:       "Add filetype to list to ignore (ex: '.hpp')",
				Destination: &setFlags.new_blacklist_item,
			},
		},
		Action: func(cCtx *cli.Context) error {
			var path string
			if cCtx.NArg() > 0 {
				path = cCtx.Args().Get(0)
			} else {
				fmt.Println("Error: No filepath entered")
			}

			count, file_count := count_through_directory(path, setFlags)
			output_value(count, file_count)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
