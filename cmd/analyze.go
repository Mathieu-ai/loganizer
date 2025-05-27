package cmd

import (
	"fmt"
	"os"

	analyzer "loganizer/internal/analyser"
	"loganizer/internal/config"
	"loganizer/internal/reporter"

	"github.com/spf13/cobra"
)

var (
	configPath string
	outputPath string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze log files based on configuration",
	Long: `Analyze multiple log files concurrently based on a JSON configuration file.
The command will process each log file in parallel and generate a report.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			fmt.Println("Error: config file path is required")
			os.Exit(1)
		}

		// Load configuration
		configs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			os.Exit(1)
		}

		if len(configs) == 0 {
			fmt.Println("No log configurations found")
			return
		}

		fmt.Printf("Starting analysis of %d log files...\n", len(configs))

		// Analyze logs
		results := analyzer.AnalyzeLogs(configs)

		// Print results
		reporter.PrintResults(results)

		// Export to JSON if output path is provided
		if outputPath != "" {
			if err := reporter.ExportResults(results, outputPath); err != nil {
				fmt.Printf("Error exporting results: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("\nResults exported to: %s\n", outputPath)
		}

		// Summary
		successful := 0
		failed := 0
		for _, result := range results {
			if result.Status == "OK" {
				successful++
			} else {
				failed++
			}
		}

		fmt.Printf("\n=== Summary ===\n")
		fmt.Printf("Total logs analyzed: %d\n", len(results))
		fmt.Printf("Successful: %d\n", successful)
		fmt.Printf("Failed: %d\n", failed)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to configuration JSON file (required)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to output JSON file (optional)")
	analyzeCmd.MarkFlagRequired("config")
}
