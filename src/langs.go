package src

type t_langs map[string][]string

// Key: lowercase language name based off the file extension
// values: list of strings: single comment, block open, block close

var LANGUAGES_COMMENTS t_langs = t_langs{
	"c":   []string{"//", "/*", "*/"},
	"cc":  []string{"//", "/*", "*/"},
	"cpp": []string{"//", "/*", "*/"},
	"h":   []string{"//", "/*", "*/"},
	"hpp": []string{"//", "/*", "*/"},
	"go":  []string{"//", "/*", "*/"},
	"py":  []string{"#"},                              // python
	"sh":  []string{"#", ": << 'COMMENT'", "COMMENT"}, //bash

}
