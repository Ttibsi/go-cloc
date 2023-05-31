package cmd

import (
	"os"
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Setup() {
	rootCmd.AddCommand(gitRepoCmd)
	rootCmd.AddCommand(findCmd)

	rootCmd.Flags().StringP("exclude-dir", "", "", "Directory to ignore when searching")
	rootCmd.Flags().StringP("exclude", "", "", "Exclude all files that match the given regex")
	rootCmd.Flags().BoolP("ignore-whitespace", "i", false, "Ignore empty lines")
	rootCmd.Flags().BoolP("ignore-comments", "", false, "Ignore comment lines")
	rootCmd.Flags().BoolP("no-recurse", "", false, "Read the files in the current directory, but don't go below")
	rootCmd.Flags().BoolP("version", "v", false, "Print version info and exit")

	findCmd.Flags().StringP("lang", "l", "", "Language to scan for")
}
