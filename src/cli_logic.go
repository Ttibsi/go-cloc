package src

import (
	"os"
)

type Flags struct {
	ExcludeDir       []string
	Exclude          []string
	IgnoreWhitespace bool
	IgnoreComments   bool
	NoRecurse        bool
	Version          bool
	FindLang         string
	GitOnly          bool
	Directory        []string
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Setup() {
	// rootCmd.AddCommand(gitRepoCmd)
	// rootCmd.AddCommand(findCmd)

	var ExcludeDirValues []string
	var ExcludeRegexValues []string

	rootCmd.Flags().
		StringSliceVarP(&ExcludeDirValues, "exclude-dir", "", []string{}, "Directory to ignore when searching")
	rootCmd.Flags().
		StringSliceVarP(&ExcludeRegexValues, "exclude", "", []string{}, "Exclude all files that match the given regex")
	rootCmd.Flags().BoolP("ignore-whitespace", "i", false, "Ignore empty lines")
	rootCmd.Flags().BoolP("ignore-comments", "", false, "Ignore comment lines")
	rootCmd.Flags().
		BoolP("no-recurse", "", false, "Read the files in the current directory, but don't go below")
	rootCmd.Flags().BoolP("version", "v", false, "Print version info and exit")
	rootCmd.Flags().StringP("lang", "l", "", "Language to scan for")
	rootCmd.Flags().BoolP("git-repo", "g", false, "Scan the entire git repository")
}
