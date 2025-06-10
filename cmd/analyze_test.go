package cmd

import (
	"bytes"
	"loganizer/internal/reporter"
	"os"
	"testing"
)

func TestFilterResultsByStatus(t *testing.T) {
	results := []reporter.LogResult{
		{LogID: "1", Status: "OK"},
		{LogID: "2", Status: "FAILED"},
	}
	ok := filterResultsByStatus(results, "OK")
	if len(ok) != 1 || ok[0].LogID != "1" {
		t.Errorf("Expected only OK result, got %+v", ok)
	}
}

func TestPrintSummary(t *testing.T) {
	results := []reporter.LogResult{
		{LogID: "1", Status: "OK"},
		{LogID: "2", Status: "FAILED", Message: "fail"},
	}
	// Capture output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	printSummary(results)
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	out := buf.String()
	if out == "" {
		t.Error("Expected output from printSummary")
	}
}
