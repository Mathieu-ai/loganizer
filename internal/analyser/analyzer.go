package analyzer

import (
	"errors"
	"fmt"
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
