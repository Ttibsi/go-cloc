package src

type t_langs map[string][]string

// Key: lowercase language name based off the file extension
// values: list of strings: single comment, block open, block close

var LANGUAGES_COMMENTS t_langs = t_langs{
	"":      []string{""},
	".c":    []string{"//", "/*", "*/"},
	".cc":   []string{"//", "/*", "*/"},
	".cfg":  []string{"#"},
	".cpp":  []string{"//", "/*", "*/"},
	".h":    []string{"//", "/*", "*/"},
	".hpp":  []string{"//", "/*", "*/"},
	".go":   []string{"//", "/*", "*/"},
	".ini":  []string{""},
	".json": []string{""},
	".md":   []string{"", "<!--", "-->"},
	".mod":  []string{"//", "/*", "*/"},
	".py":   []string{"#"},                              // python
	".sh":   []string{"#", ": << 'COMMENT'", "COMMENT"}, //bash
	".txt":  []string{""},
	".sum":  []string{"//", "/*", "*/"},
	".yaml": []string{"#"},
	".yml":  []string{"#"},
}

var BLACKLIST []string = []string{
	".DS_Store",
	".pyc",
	".png",
	".jpg",
	".gif",
	".bmp",
}
