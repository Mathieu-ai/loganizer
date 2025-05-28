package analyzer

import (
	"errors"
	"fmt"
)

// FileNotFoundError represents file system access errors during log analysis
type FileNotFoundError struct {
	Path string
}

// Error implements the error interface for FileNotFoundError
//
// Returns the formatted error message string.
func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s", e.Path)
}

// ParsingError represents log parsing failures with context
type ParsingError struct {
	LogID   string
	Message string
}

// Error implements the error interface for ParsingError
//
// Returns the formatted error message string with log ID context.
func (e *ParsingError) Error() string {
	return fmt.Sprintf("parsing error for log %s: %s", e.LogID, e.Message)
}

// IsFileNotFoundError checks error type using errors.Is pattern
//
// Parameters:
//   - err: The error to check
//
// Returns true if error is of type FileNotFoundError.
func IsFileNotFoundError(err error) bool {
	var fnfErr *FileNotFoundError
	return errors.As(err, &fnfErr)
}

// IsParsingError checks error type using errors.As pattern
//
// Parameters:
//   - err: The error to check
//
// Returns true if error is of type ParsingError.
func IsParsingError(err error) bool {
	var parseErr *ParsingError
	return errors.As(err, &parseErr)
}

// GetFileNotFoundError extracts FileNotFoundError from error chain
//
// Parameters:
//   - err: The error to extract from
//
// Returns FileNotFoundError instance and boolean indicating success.
func GetFileNotFoundError(err error) (*FileNotFoundError, bool) {
	var fnfErr *FileNotFoundError
	if errors.As(err, &fnfErr) {
		return fnfErr, true
	}
	return nil, false
}

// GetParsingError extracts ParsingError from error chain
//
// Parameters:
//   - err: The error to extract from
//
// Returns ParsingError instance and boolean indicating success.
func GetParsingError(err error) (*ParsingError, bool) {
	var parseErr *ParsingError
	if errors.As(err, &parseErr) {
		return parseErr, true
	}
	return nil, false
}
