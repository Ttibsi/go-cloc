package src

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// commands
var rootCmd = &cobra.Command{
	Use:   "go-cloc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		flags := Flags{}

		if len(args) == 0 {
			// PWD
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err.Error())
			}

			flags.Directory = append(flags.Directory, pwd)
		} else {
			for _, val := range args {
				flags.Directory = append(flags.Directory, val)
			}
		}

		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Println(err.Error())
		}
		if version {
			flags.Version = true
		}

		excluded, err := cmd.Flags().GetStringSlice("exclude-dir")
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(excluded) != 0 {
			flags.ExcludeDir = excluded
		}

		excluded_regex, err := cmd.Flags().GetStringSlice("exclude")
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(excluded_regex) != 0 {
			flags.Exclude = excluded_regex
		}

		ignoreWhitespace, err := cmd.Flags().GetBool("ignore-whitespace")
		if err != nil {
			fmt.Println(err.Error())
		}
		if ignoreWhitespace {
			flags.IgnoreWhitespace = true
		}

		ignoreComments, err := cmd.Flags().GetBool("ignore-comments")
		if err != nil {
			fmt.Println(err.Error())
		}
		if ignoreComments {
			flags.IgnoreComments = true
		}

		noRecurse, err := cmd.Flags().GetBool("no-recurse")
		if err != nil {
			fmt.Println(err.Error())
		}
		if noRecurse {
			flags.NoRecurse = true
		}

		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		if lang != "" {
			flags.FindLang = lang
		}

		git, err := cmd.Flags().GetBool("git-repo")
		if err != nil {
			fmt.Println(err.Error())
		}

		if git {
			flags.GitOnly = true
		}

		StartLogic(flags)
	},
}

// var gitRepoCmd = &cobra.Command{
// 	Use: "git-repo",
// 	Short: "Searches through the entire git repository you're currently in",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("gitRepoCmd")
// 	},
// }
//
// var findCmd = &cobra.Command{
// 	Use: "find",
// 	Short: "Find all files of a specified language. --lang flag required",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		lang, err := cmd.Flags().GetString("lang")
// 		if err != nil {
// 			os.Exit(1)
// 		}
//
// 		if lang != "" {
// 			fmt.Printf("Searching for: %v\n", lang)
// 		} else {
// 			fmt.Println("findCmd")
// 		}
// 	},
// }

func init() { Setup() }
