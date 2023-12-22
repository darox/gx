package main

import (
	"os"

	"github.com/darox/gx/pkg/gx"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "gx -t <path> -s <path> -u <url> -b <branch>",
	Long: "gx is a tool for selectively downloading files and folders from git repositories",
	Run: func(cmd *cobra.Command, args []string) {

		target := cmd.Flag("target").Value.String()
		source := cmd.Flag("source").Value.String()
		url := cmd.Flag("url").Value.String()
		branch := cmd.Flag("branch").Value.String()

		repository, err := gx.NewRepository(url, branch, source)
		if err != nil {
			os.Exit(1)
		}

		repository.Clone()

		e, err := repository.Extract()

		if err != nil {
			os.Exit(1)
		}

		err = e.Write(target)

		if err != nil {
			os.Exit(1)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("branch", "b", "", "the branch name to extract from")
	rootCmd.PersistentFlags().StringP("url", "u", "", "the url of the git repository")
	rootCmd.PersistentFlags().StringP("target", "t", "", "the target path to write to")
	rootCmd.PersistentFlags().StringP("source", "s", "", "the source path to extract from")
}
