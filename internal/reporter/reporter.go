package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogResult represents the result of analyzing a single log
type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

// ExportResults exports results to a JSON file with automatic directory creation
func ExportResults(results []LogResult, outputPath string) error {
	// Create directories if they don't exist (BONUS feature)
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

// PrintResults prints results to console with detailed formatting
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

// GenerateTimestampedFilename generates a filename with current date (BONUS feature)
func GenerateTimestampedFilename(basePath string) string {
	now := time.Now()
	timestamp := now.Format("060102") // YYMMDD format

	dir := filepath.Dir(basePath)
	ext := filepath.Ext(basePath)
	name := filepath.Base(basePath)
	nameWithoutExt := name[:len(name)-len(ext)]

	timestampedName := fmt.Sprintf("%s_%s%s", timestamp, nameWithoutExt, ext)
	return filepath.Join(dir, timestampedName)
}
