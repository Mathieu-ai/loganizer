package analyzer

import (
	"errors"
	"fmt"
)

// FileNotFoundError represents a file that is not found or inaccessible error
type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

// ParsingError represents a parsing error during log analysis
type ParsingError struct {
	LogID   string
	Message string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("parsing error for log %s: %s", e.LogID, e.Message)
}

// IsFileNotFoundError checks if the error is a FileNotFoundError using errors.Is
func IsFileNotFoundError(err error) bool {
	var fnfErr *FileNotFoundError
	return errors.As(err, &fnfErr)
}

// IsParsingError checks if the error is a ParsingError using errors.As
func IsParsingError(err error) bool {
	var parseErr *ParsingError
	return errors.As(err, &parseErr)
}

// GetFileNotFoundError extracts FileNotFoundError details if present
func GetFileNotFoundError(err error) (*FileNotFoundError, bool) {
	var fnfErr *FileNotFoundError
	if errors.As(err, &fnfErr) {
		return fnfErr, true
	}
	return nil, false
}

// GetParsingError extracts ParsingError details if present
func GetParsingError(err error) (*ParsingError, bool) {
	var parseErr *ParsingError
	if errors.As(err, &parseErr) {
		return parseErr, true
	}
	return nil, false
}
