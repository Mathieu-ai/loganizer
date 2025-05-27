package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "A CLI tool for analyzing log files",
	Long:  "LogAnalyzer is a command-line tool that helps system administrators analyze log files from various sources in parallel with robust error handling.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
