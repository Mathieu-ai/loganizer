package analyzer

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"loganizer/internal/config"
	"loganizer/internal/reporter"
)

func AnalyzeLog(logConfig config.LogConfig, results chan<- reporter.LogResult, wg *sync.WaitGroup) {
	defer wg.Done()

	result := reporter.LogResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	// Check if file exists and is readable
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

	// Simulate analysis with random sleep (50-200ms as per README)
	sleepDuration := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(sleepDuration)

	// Simulate random parsing error (10% chance as per README)
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

// AnalyzeLogs analyzes multiple log files concurrently
func AnalyzeLogs(configs []config.LogConfig) []reporter.LogResult {
	var wg sync.WaitGroup
	results := make(chan reporter.LogResult, len(configs))

	// Launch goroutines for each log
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
