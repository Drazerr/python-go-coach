package main

import (
	"strings"
	"testing"
)

func TestSummarize(t *testing.T) {
	input := `1.1.1.1 - - [08/Jul/2026:22:14:01 +0000] "GET /a HTTP/1.1" 500 10
1.1.1.1 - - [08/Jul/2026:22:14:02 +0000] "GET /b HTTP/1.1" 200 10
garbage`
	got, err := summarize(strings.NewReader(input))
	if err != nil {
		t.Fatalf("summarize returned error: %v", err)
	}
	if got.TotalLines != 3 {
		t.Errorf("TotalLines = %d, want 3", got.TotalLines)
	}
	// your turn: TotalErrors, Malformed, ErrorCounts["500"], ErrorURLs["/a"]
	if got.TotalErrors != 1 {
		t.Errorf("TotalErrors = %d, want 1", got.TotalErrors)
	}
	if got.Malformed != 1 {
		t.Errorf("Malformed = %d, want 1", got.Malformed)
	}
	if got.ErrorCounts["500"] != 1 {
		t.Errorf("Malformed = %d, want 1", got.ErrorCounts["500"])
	}
	if got.ErrorURLs["/a"] != 1 {
		t.Errorf("Malformed = %d, want 1", got.ErrorURLs["/a"])
	}
}
