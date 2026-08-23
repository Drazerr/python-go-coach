# python-go-lab

A learning repo where I'm refreshing my Python scripting skills while learning Go.

Rather than working through syntax-only exercises, I'm using small automation, DevOps, and SRE-style problems. Most exercises are implemented in both Python and Go so I can compare how the two languages approach the same problem.

The goal is simple: write more code, test it, understand the differences between the languages, and gradually move toward more realistic tooling.

## What I'm focusing on

* Python scripting and automation
* Learning idiomatic Go
* CLI utilities
* File and configuration processing
* HTTP and API interactions
* Error handling
* Testing
* JSON, YAML, and CSV
* Automation for infrastructure and operations
* Concurrency and resilient tooling as the exercises get more advanced

## How the repo is organized

Each exercise gets its own session directory.

Inside the session, I keep the Python and Go implementations separate so I can solve the same problem in both languages.

```text
.
├── session-01-log-parser/
│   ├── README.md
│   ├── python/
│   └── go/
│
└── README.md
```

## Sessions

| Session | Exercise                             | Main focus                                                 |
| ------- | ------------------------------------ | ---------------------------------------------------------- |
| 01      | [Log Parser](session-01-log-parser/) | File processing, error handling, data aggregation, testing |

I'll add new sessions here as I work through them.

## Why Python and Go?

Python is already familiar to me, which makes it a useful reference point while learning Go.

Solving the same problem twice lets me compare things such as:

```text
data structures
error handling
testing
CLI design
standard library usage
code organization
performance considerations
```

I'm less interested in translating Python line-for-line into Go. The goal is to understand how I would naturally solve the problem in each language.

## Running the exercises

Each session has its own README with the problem, implementation notes, and commands for running and testing both versions.

For example:

```bash
cd session-01-log-parser
```

Then follow the instructions in that session's README.

## About this repo

This is intentionally a learning repository.

Some implementations will probably change as I learn better patterns or find more idiomatic ways to write the same code. I'm keeping that progression visible rather than trying to make every exercise look like a finished production project.
