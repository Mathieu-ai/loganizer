package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogResult represents the analysis outcome for a single log file
type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

/**
 * ExportResults serializes results to JSON file with directory auto-creation.
 * Creates intermediate directories if they don't exist.
 *
 * @param results Slice of LogResult to export
 * @param outputPath Path where the JSON file should be written
 * @return Error if file creation or JSON encoding fails
 */
func ExportResults(results []LogResult, outputPath string) error {
	// Ensure output directory exists
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

/**
 * PrintResults outputs formatted analysis results to stdout.
 * Displays each result with ID, path, status, message, and error details.
 *
 * @param results Slice of LogResult to display
 */
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

/**
 * GenerateTimestampedFilename prepends YYMMDD timestamp to base filename.
 * Preserves directory structure and file extension.
 *
 * @param basePath Original file path to timestamp
 * @return New file path with timestamp prefix in YYMMDD format
 */
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
