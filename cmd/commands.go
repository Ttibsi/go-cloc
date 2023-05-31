package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//commands
var rootCmd = &cobra.Command{
	Use:   "go-cloc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) { 
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if version {
			fmt.Println("Version: ") 
			os.Exit(0)
		}

		excludedDir, err := cmd.Flags().GetString("exclude-dir")
		if err != nil {
			os.Exit(1)
		}

		if excludedDir != "" {
			fmt.Printf("Excluding: %v\n", excludedDir)
		}

		excluded, err := cmd.Flags().GetString("exclude")
		if err != nil {
			os.Exit(1)
		}

		if excluded != "" {
			fmt.Printf("Excluding: %v\n", excluded)
		}

		ignoreWhitespace, err := cmd.Flags().GetBool("ignore-whitespace")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if ignoreWhitespace {
			fmt.Println("ignore whitespace") 
		}

		ignoreComments, err := cmd.Flags().GetBool("ignore-comments")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if ignoreComments {
			fmt.Println("ignore comments") 
		}

		noRecurse, err := cmd.Flags().GetBool("no-recurse")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		if noRecurse {
			fmt.Println("no recurse") 
		}

		fmt.Println("root") 
	},
}

var gitRepoCmd = &cobra.Command{
	Use: "git-repo",
	Short: "Searches through the entire git repository you're currently in",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitRepoCmd")
	},
}

var findCmd = &cobra.Command{
	Use: "find",
	Short: "Find all files of a specified language. --lang flag required",
	Run: func(cmd *cobra.Command, args []string) {
		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		if lang != "" {
			fmt.Printf("Searching for: %v\n", lang)
		} else {
			fmt.Println("findCmd") 
		}
	},
}

func init() { Setup() }
