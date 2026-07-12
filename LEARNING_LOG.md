# LEARNING_LOG.md

## Meta
- Session count: 1 (IN PROGRESS — checkpoint, resume in new chat as Session 1)
- Last updated: 2026-07-12
- Environment: macOS, Python 3.14.6, Go 1.26.4, 60 min/session
- Repo: https://github.com/Drazerr/python-go-coach (push pending)

## Baseline Assessment (2026-07-09)
- Python: 1.5→2 (performed above baseline in S1) | Go: 0→1
- Linux/shell: 3 | Git: 2 | Testing: 2 | HTTP/APIs: 0 (priority gap, S4)
- Docker: 3 | AWS: 3 | Azure: 2 | GCP: 1 | K8s: 2
- Target: DevOps/SRE/Infra/Platform. Ops-strong, code-light. Ops contexts everywhere.

## Roadmap Summary (adaptive)
- S1: Log error summarizer ✅ (code done, closing steps pending)
- S2: Config/env validator — S3: JSON↔YAML — S4: HTTP health checker — S5: Review #1
- S6–9: CSV reports, cert expiry, structured logging, CLI subcommands — S10: Review #2
- S11–14: Stage 2 (pagination, retries, stale resources, config diff; branch-per-exercise Git)
- S15: Review #3 — S16–19: SQLite, DI/interfaces, resilient client, first concurrency — S20: Review #4
- S21–24: Stage 3 (worker pools, concurrent checker, log pipeline, K8s auditor) — S25: Review #5
- S26–30: Capstone + final mock interview
- Adaptation: S1 pace was strong (Python above baseline; Go picked up fast after step-down).
  If S2–3 also land easily, compress Stage 1 by merging S3 concepts into S2 or S4.

## Session 1 (2026-07-12) — Log error summarizer — CHECKPOINT
**Done:** Python v1 + dataclass refactor + pytest (passing, verified). Go version complete:
struct, io.Reader seam, (LogSummary, error) return, scanner.Err(), sorted output via helper,
go test passing (verified, incl. deliberate-failure drill). Comparison discussion delivered.
**Remaining to close S1 (do FIRST on resume, in the same order):**
1. Answers to 3 interview-reflection questions (explain summarize; 10GB log; add per-IP counting)
2. Two commit messages for feedback; repo created/pushed; go.mod path fixed
   (was github.com/Drazerr/python-go-coach/tree/main/... — 'tree/main' must go; folder nesting was doubled)
3. Coach confirms → final S1 log

## Progress
- Python concepts: file iteration, dict.get counting, argparse, dataclasses, OSError handling,
  exit codes, pytest basics, testable-function design (iterable seam)
- Go concepts: modules, os.Args, os.Open + if err != nil, defer, bufio.Scanner + Err(),
  strings.Fields, strconv.Atoi, maps (make, ++), structs (named-field literals),
  multi-value/error returns, sort.Slice, strings.NewReader, table-stakes testing, go vet via go test
- Common mistakes: Fprintln vs Println(os.Stderr,...); int vs string map key (caught by compiler);
  scanner.Err() checked inside loop (dead code); capitalized locals; positional struct literal;
  copy-paste test assertions (wrong messages, duplicated assert, one target untested);
  pasted non-compiling code without running it (coached: compiler first, paste real output)
- Topics needing revision: sort tie-breaking; %-verbs vs types (vet); test message hygiene
- Testing: both languages, verified runs, watched-it-fail drill done. Claim of 2 looks accurate.
- Debugging: reads compiler errors with guidance; not yet independently
- Interview strengths: made code testable without prompting (seam design in both languages)
- Interview weaknesses: articulating comparisons (bullets too thin — coach expanded);
  explaining code aloud untested (Q1 of reflection targets this)

## Current State
- Difficulty: Easy (S1 suggests Easy-medium may come early; decide after S2)
- Rubric (S1, provisional — finalize at close): Correctness: Competent | Robustness: Developing
  | Readability: Competent | Structure: Competent | Testing: Developing | Error handling:
  Competent | Language usage: Developing (Go) / Competent (Py) | Operational thinking:
  Developing | Interview communication: Needs work→Developing (pending Q1–3)
- Revision exercise before S2: add alphabetical tie-breaker to printSortedCounts so equal
  counts print in stable order (observed real tie-instability in run 4); rerun go test

## Recommended Next Session
**Close S1 checkpoint first (items above), then S2 — Config/env validator (Easy).**
Tool reads required env vars / a config file, validates presence + types + simple rules,
reports all problems (not just the first), exits non-zero on failure. Python: dataclass
config, pytest parametrize. Go: struct validation, error wrapping with fmt.Errorf %w,
table-driven tests — the canonical Go test pattern, building on S1's single test.
