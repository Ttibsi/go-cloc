package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func path_trawl(dir string) []string {
	all_paths := make([]string, 1000) // slice of max length 1000

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			} else {

				for _, elem := range strings.Split(path, "\n") {
					//TODO: skip dot files
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

func file_length(file string) int {

}

func main() {
	PLACEHOLDER := "/users/ashisbitt/workspace/gh-stats"
	dir := PLACEHOLDER
	var count int = 0

	all_files := path_trawl(dir)

	for _, file := range all_files {
		count += file_length(file)
	}

	fmt.Println(count)

}
