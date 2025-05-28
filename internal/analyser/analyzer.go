package analyzer

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"loganizer/internal/config"
	"loganizer/internal/reporter"
)

/**
 * AnalyzeLog processes a single log file and sends results through channel.
 * Simulates file access validation and random parsing errors (10% probability).
 * Sleep duration: 50-200ms to simulate I/O operations.
 *
 * @param logConfig The configuration for the log file to analyze
 * @param results Channel to send analysis results to
 * @param wg WaitGroup for synchronization
 */
func AnalyzeLog(logConfig config.LogConfig, results chan<- reporter.LogResult, wg *sync.WaitGroup) {
	defer wg.Done()

	result := reporter.LogResult{
		LogID:    logConfig.ID,
		FilePath: logConfig.Path,
	}

	// Validate file accessibility
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

	// Simulate I/O latency with random sleep (50-200ms)
	sleepDuration := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(sleepDuration)

	// Simulate parsing failure with 10% probability
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

/**
 * AnalyzeLogs orchestrates concurrent log analysis using goroutines and channels.
 * Returns aggregated results from all worker goroutines.
 *
 * @param configs Slice of log configurations to analyze
 * @return Slice of analysis results from all processed logs
 */
func AnalyzeLogs(configs []config.LogConfig) []reporter.LogResult {
	var wg sync.WaitGroup
	results := make(chan reporter.LogResult, len(configs))

	// Spawn worker goroutines for concurrent processing
	for _, cfg := range configs {
		wg.Add(1)
		go AnalyzeLog(cfg, results, &wg)
	}

	// Wait for all workers to complete
	wg.Wait()
	close(results)

	// Aggregate results from channel
	var logResults []reporter.LogResult
	for result := range results {
		logResults = append(logResults, result)
	}

	return logResults
}
