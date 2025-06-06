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
	configPath   string
	outputPath   string
	statusFilter string
	useTimestamp bool
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze log files based on configuration",
	Long: `Analyze multiple log files concurrently based on a JSON configuration file.
The command will process each log file in parallel using goroutines and generate a comprehensive report.
Supports filtering by status and timestamped output files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			fmt.Println("Error: config file path is required")
			os.Exit(1)
		}

		// Load configuration from JSON
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

		// Execute concurrent analysis with goroutines
		results := analyzer.AnalyzeLogs(configs)

		// Apply status filtering if specified
		if statusFilter != "" {
			results = filterResultsByStatus(results, statusFilter)
		}

		// Terminal output
		reporter.PrintResults(results)

		// JSON export with optional timestamping
		if outputPath != "" {
			finalOutputPath := outputPath

			// Apply timestamp prefix to filename
			if useTimestamp {
				finalOutputPath = reporter.GenerateTimestampedFilename(outputPath)
			}

			if err := reporter.ExportResults(results, finalOutputPath); err != nil {
				fmt.Printf("Error exporting results: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("\nResults exported to: %s\n", finalOutputPath)
		}

		// Output analysis metrics
		printSummary(results)
	},
}

/**
 * filterResultsByStatus filters LogResult slice by status field.
 *
 * @param results Slice of LogResult to filter
 * @param status Status value to filter by ("OK" or "FAILED")
 * @return Filtered slice containing only results with matching status
 */
func filterResultsByStatus(results []reporter.LogResult, status string) []reporter.LogResult {
	var filtered []reporter.LogResult
	for _, result := range results {
		if result.Status == status {
			filtered = append(filtered, result)
		}
	}
	return filtered
}

/**
 * printSummary outputs analysis metrics and failure details.
 * Displays total count, success/failure breakdown, and detailed error information.
 *
 * @param results Slice of LogResult to summarize
 */
func printSummary(results []reporter.LogResult) {
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

	if failed > 0 {
		fmt.Printf("\nFailed logs breakdown:\n")
		for _, result := range results {
			if result.Status == "FAILED" {
				fmt.Printf("- %s: %s\n", result.LogID, result.Message)
			}
		}
	}
}

/**
 * init registers the analyze command with flags.
 * Sets up command-line flags for config path, output path, status filter, and timestamp options.
 */
func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to configuration JSON file (required)")
	analyzeCmd.MarkFlagRequired("config")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to output JSON file (optional)")
	analyzeCmd.Flags().StringVar(&statusFilter, "status", "", "Filter results by status (OK or FAILED)")
	analyzeCmd.Flags().BoolVar(&useTimestamp, "timestamp", false, "Add timestamp to output filename")
}
