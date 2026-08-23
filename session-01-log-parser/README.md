# Session 01: Log Parser

For the first exercise, I built a small command-line tool that reads a web server log and summarizes HTTP errors.

I implemented the same basic behavior in Python and Go so I could compare file processing, data structures, error handling, and testing in both languages.

## Problem

Given a log file, the program should report:

* total lines processed
* malformed lines skipped
* total HTTP errors
* errors grouped by status code
* errors grouped by URL

For this exercise, an HTTP status code of `400` or greater is treated as an error.

## Structure

```text
session-01-log-parser/
├── README.md
│
├── python/
│   ├── log_summary.py
│   ├── test_log_summary.py
│   └── sample.log
│
└── go/
    ├── go.mod
    ├── main.go
    ├── main_test.go
    └── sample.log
```

## Python

The Python implementation reads the log as an iterable and returns the results in a `LogSummary` dataclass.

Some of the things I practiced here:

* `argparse`
* dataclasses
* dictionaries
* file iteration
* exception handling
* exit codes
* separating parsing from output
* pytest

### Run

```bash
cd python
python3 log_summary.py sample.log
```

### Test

```bash
pytest
```

## Go

The Go implementation solves the same problem using Go's standard library.

Instead of tying the parser directly to a file, the summarizer accepts an `io.Reader`. That makes it possible to use a file in the CLI while passing an in-memory reader from tests.

Some of the Go concepts I practiced:

* structs
* maps
* `io.Reader`
* `bufio.Scanner`
* `strconv.Atoi`
* explicit error returns
* sorting map output
* Go modules
* `go test`

### Run

```bash
cd go
go run . sample.log
```

### Test

```bash
go test ./...
```

## Comparing the two

The core logic is very similar, but the languages encourage slightly different approaches.

Python makes the aggregation code compact, especially when working with dictionaries and a dataclass for the result.

Go requires more explicit handling of parsing and errors, but using interfaces such as `io.Reader` makes the boundary between input and processing very clear.

Both versions keep the log-processing logic separate from the command-line entry point, which also makes the code easier to test.

## What I took away

The main lesson from this session wasn't log parsing itself.

It was seeing how the same small operational tool looks in two languages:

```text
Python
  quick to express the logic
  lightweight data handling
  concise error handling

Go
  explicit control flow
  explicit error handling
  strong standard-library patterns
  interfaces make test boundaries clear
```

I'll use the same side-by-side approach for the next exercises and gradually add more realistic problems.
