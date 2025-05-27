package analyzer

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"loganizer/internal/config"
	"loganizer/internal/reporter"
)

// FileNotFoundError represents a file that is not found error
type FileNotFoundError struct {
	Path string
}

func (e FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

// ParsingError represents the parsing error
type ParsingError struct {
	LogID   string
	Message string
}

func (e ParsingError) Error() string {
	return fmt.Sprintf("parsing error for log %s: %s", e.LogID, e.Message)
}

// IsFileNotFoundError checks if the error refers to FileNotFoundError
func IsFileNotFoundError(err error) bool {
	var fnfErr FileNotFoundError
	return errors.As(err, &fnfErr)
}

// IsParsingError checks if the error refers to ParsingError
func IsParsingError(err error) bool {
	var parseErr ParsingError
	return errors.As(parseErr, &parseErr)
}

func AnalyzeLog(logConfig config.LogConfig, results chan<- reporter.LogResult, wg *sync.WaitGroup) {
	defer wg.Done()

	result := reporter.LogResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	// Check if file exists
	if _, err := os.Stat(logConfig.Path); err != nil {
		if os.IsNotExist(err) {
			fileErr := &FileNotFoundError{Path: logConfig.Path}
			result.Status = "FAILED"
			result.Message = "Fichier introuvable."
			result.ErrorDetails = fileErr.Error()
		} else {
			result.Status = "FAILED"
			result.Message = "Fichier inaccessible."
			result.ErrorDetails = err.Error()
		}
		results <- result
		return
	}

	// Simulate analysis with random sleep
	sleepDuration := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(sleepDuration)

	// Simulate random parsing error
	if rand.Float32() < 0.1 {
		parseErr := &ParsingError{
			LogID:   logConfig.ID,
			Message: "erreur de format de ligne",
		}
		result.Status = "FAILED"
		result.Message = "Erreur de parsing."
		result.ErrorDetails = parseErr.Error()
	} else {
		result.Status = "OK"
		result.Message = "Analyse terminée avec succès."
		result.ErrorDetails = ""
	}

	results <- result
}

// analyzes multiple log files concurrently
func AnalyzeLogs(configs []config.LogConfig) []reporter.LogResult {
	var wg sync.WaitGroup
	results := make(chan reporter.LogResult, len(configs))

	// goroutines for each log
	for _, cfg := range configs {
		wg.Add(1)
		go AnalyzeLog(cfg, results, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	close(results)

	// Collect results
	var logResults []reporter.LogResult
	for result := range results {
		logResults = append(logResults, result)
	}

	return logResults
}
