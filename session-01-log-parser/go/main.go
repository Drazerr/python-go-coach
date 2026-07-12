package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LogSummary struct {
	TotalLines  int
	TotalErrors int
	Malformed   int
	ErrorCounts map[string]int
	ErrorURLs   map[string]int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: session01 <logfile>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer f.Close()
	data, err := summarize(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading log:", err)
		os.Exit(1)
	}
	printSummary(data)

}

func summarize(r io.Reader) (LogSummary, error) {

	totalLines := 0
	totalErrors := 0
	malformed := 0
	errorCounts := make(map[string]int)
	errorURLs := make(map[string]int)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		totalLines++
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 10 {
			malformed++
			continue
		}
		statusStr := parts[len(parts)-2]
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			malformed++
			continue
		}
		if status >= 400 {
			totalErrors++
			errorCounts[statusStr]++
			url := parts[6]
			errorURLs[url]++

		}
	}

	if err := scanner.Err(); err != nil {
		return LogSummary{}, err
	}

	return LogSummary{
		TotalLines:  totalLines,
		TotalErrors: totalErrors,
		Malformed:   malformed,
		ErrorCounts: errorCounts,
		ErrorURLs:   errorURLs,
	}, nil
}

func printSummary(data LogSummary) {

	fmt.Printf("Total lines: %d\n", data.TotalLines)
	fmt.Printf("Malformed lines skipped: %d\n", data.Malformed)
	fmt.Printf("Total errors (status >= 400): %d\n", data.TotalErrors)
	fmt.Println()
	fmt.Println("Errors by status code:")
	printSortedCounts(data.ErrorCounts)
	fmt.Println("Errors by URL:")
	printSortedCounts(data.ErrorURLs)

}

func printSortedCounts(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]] // by count, descending
	})
	for _, k := range keys {
		fmt.Printf("  %s: %d\n", k, m[k])
	}
}
