#!/usr/bin/env python3

import argparse
import sys
from typing import Iterable, Dict
from dataclasses import dataclass

@dataclass
class LogSummary:
    total_lines: int
    total_errors: int
    malformed: int
    error_counts: Dict[str, int]
    error_urls: Dict[str, int]

def main():
    log_file = parse_args()
    try:
        f = open(log_file, "r")
    except OSError as e:
        print(f"File system error on '{log_file}': {e}", file=sys.stderr)
        sys.exit(1)
    with f:
        data = summarize(f)
        print_output(data)


def summarize(lines: Iterable[str]) -> LogSummary:
    error_counts = {}
    total_lines = 0
    total_errors = 0
    malformed = 0
    error_urls = {}
    for line in lines:
        total_lines += 1    
        parts = line.split()
        if len(parts) < 10 or not parts[-2].isdigit():
            malformed += 1
            continue
        status = parts[-2]
        if int(status) >= 400:
            total_errors += 1
            error_counts[status] = error_counts.get(status, 0) + 1
            url = parts[6]
            error_urls[url] = error_urls.get(url, 0) + 1
    
    return LogSummary(
        total_lines=total_lines,
        total_errors=total_errors,
        malformed=malformed,
        error_counts=error_counts,
        error_urls=error_urls
    )

def print_output(log_summary: LogSummary):
    print(f"Total lines: {log_summary.total_lines}")
    print(f"Malformed lines skipped: {log_summary.malformed}")
    print(f"Total errors (status >= 400): {log_summary.total_errors}")
    print()
    print("Errors by status code:")
    for status, count in log_summary.error_counts.items():
        print(f"  {status}: {count}")
    print()
    print("Errors by URL:")
    for url, count in log_summary.error_urls.items():
        print(f"  {url}: {count}")

def parse_args():
    parser = argparse.ArgumentParser(description='Log summary')
    parser.add_argument('log_file', type=str, help='Log file to analyze')
    args = parser.parse_args()
    return args.log_file

if __name__ == "__main__":
    main()
