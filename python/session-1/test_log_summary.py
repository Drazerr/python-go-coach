from log_summary import summarize

def test_summarize_counts_errors():
    lines = [
        '1.1.1.1 - - [08/Jul/2026:22:14:01 +0000] "GET /a HTTP/1.1" 500 10',
        '1.1.1.1 - - [08/Jul/2026:22:14:02 +0000] "GET /b HTTP/1.1" 200 10',
        "garbage",
    ]
    result = summarize(lines)
    assert result.total_lines == 3
    assert result.total_errors == 1
    assert result.malformed == 1
    assert result.error_counts == {"500": 1}
    assert result.error_urls == {"/a": 1}
    