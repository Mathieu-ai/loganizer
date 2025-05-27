package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Result of analyzing a single log
type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// Exports results to a JSON file
func ExportResults(results []LogResult, outputPath string) error {
	// Create directories if they don't exist
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		return fmt.Errorf("failed to encode results: %w", err)
	}

	return nil
}

// Prints results to console
func PrintResults(results []LogResult) {
	fmt.Println("\n=== Log Analysis Results ===")
	for _, result := range results {
		fmt.Printf("ID: %s\n", result.LogID)
		fmt.Printf("Path: %s\n", result.FilePath)
		fmt.Printf("Status: %s\n", result.Status)
		fmt.Printf("Message: %s\n", result.Message)
		if result.ErrorDetails != "" {
			fmt.Printf("Error: %s\n", result.ErrorDetails)
		}
		fmt.Println("---")
	}
}
