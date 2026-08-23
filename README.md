# python-go-lab

A small learning repo where I refresh my Python scripting skills while learning Go.

I use practical automation and infrastructure-style problems rather than language-only exercises. For most sessions, I solve the same problem in both Python and Go so I can compare how each language handles things like file processing, error handling, testing, data structures, CLI design, and concurrency.

The exercises are intentionally small. The goal is to build consistency and get more comfortable writing useful scripts in both languages.

## What I'm focusing on

* Python scripting and automation
* Learning idiomatic Go
* CLI utilities
* File and configuration processing
* HTTP and API interactions
* Error handling
* Testing
* Structured data such as JSON, YAML, and CSV
* Practical DevOps and SRE-style problems
* Eventually, concurrency and more resilient tooling

## Repository structure

Each session contains the same exercise implemented separately in Python and Go.

```text
.
├── session-01-log-parser/
│   ├── python/
│   │   ├── log_summary.py
│   │   ├── test_log_summary.py
│   │   └── sample.log
│   │
│   └── go/
│       ├── go.mod
│       ├── main.go
│       ├── main_test.go
│       └── sample.log
│
└── README.md
```

As I add more exercises, I'll keep the same general structure:

```text
session-02-...
├── python/
└── go/
```

## Session 01: Log parser

The first exercise is a simple web log summarizer.

It reads a log file and reports:

* total lines
* malformed lines
* HTTP errors
* error counts by status code
* error counts by URL

I implemented the same basic behavior in both languages and added tests for each version.

### Python

The Python version gave me a chance to revisit:

* `argparse`
* dataclasses
* dictionaries
* iterable-based functions
* file handling
* exception handling
* pytest

Run it with:

```bash
cd session-01-log-parser/python
python3 log_summary.py sample.log
```

Run the tests:

```bash
pytest
```

### Go

The Go version focuses on similar behavior using Go's standard library.

Some of the concepts I'm practicing:

* structs
* maps
* `io.Reader`
* `bufio.Scanner`
* explicit error handling
* sorting
* Go modules
* table-oriented thinking
* `go test`

Run it with:

```bash
cd session-01-log-parser/go
go run . sample.log
```

Run the tests:

```bash
go test ./...
```

## Why solve the same problem twice?

That's the main idea behind this repo.

Python is already familiar to me, so it gives me a useful reference point while learning Go.

Instead of learning Go syntax in isolation, I can compare questions such as:

```text
How would I represent this data in Python?
How would I represent it in Go?

How does each language handle errors?

How do I make the code testable?

What belongs in a function versus the CLI layer?

How does the standard library differ?
```

That comparison has been more useful to me than working through unrelated language exercises.

## Planned exercises

I'm planning to work through small problems such as:

* configuration and environment validation
* JSON and YAML conversion
* HTTP health checking
* CSV reporting
* certificate expiry checking
* structured logging
* CLI subcommands
* API pagination and retries
* stale resource detection
* configuration diffing
* SQLite
* concurrent checks in Go

The list will probably change as I go.

## Notes

This is a learning repository, so some implementations may change as I find cleaner or more idiomatic ways to solve the same problem.

The goal isn't to make every exercise production-ready. It's to practice writing code, testing it, understanding the trade-offs, and improving from one session to the next.
