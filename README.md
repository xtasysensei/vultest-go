# VulTest

VulTest is a fast and extensible CLI vulnerability scanner written in Go, designed for security researchers, penetration testers, and developers. It supports scanning web applications for vulnerabilities like XSS and SQL Injection (SQLi) using customizable depth and user-agent rotation.

## Features

- Crawl target websites and test for XSS or SQLi
- Simple CLI interface built with Cobra
- Random user-agent rotation for evasive scanning
- Built with extensibility and security in mind
- Support for concurrency using Go's goroutines and WaitGroups

## Installation

### From Source

```bash
git clone https://github.com/xtasysensei/vultest.git
cd vultest
go build -o vultest
```

### From Release (recommended)

Download the prebuilt binary from the [Releases](https://github.com/xtasysensei/vultest/releases) page for your OS and architecture, then give it execute permissions:

```bash
chmod +x vultest
./vultest --help
```

## Usage

```bash
./vultest scan --type [xss|sqli] --url http://example.com --depth [number]
```

### Options

| Flag         | Description                             | Required |
|--------------|-----------------------------------------|----------|
| `--type`     | Scan type: `xss` or `sqli`              | Yes      |
| `--url`      | Target URL to scan                      | Yes      |
| `--depth`    | Crawling depth (how deep to follow links) | No (default: 1) |

## Examples

```bash
# Scan for XSS vulnerabilities
./vultest scan --type xss --url "http://testphp.vulnweb.com" --depth 2

# Scan for SQLi vulnerabilities
./vultest scan --type sqli --url "http://example.com"
```

## Project Structure

```
vultest/
├── cmd/              # CLI commands
├── handlers/         # Scanning logic (XSS, SQLi, etc.)
├── main.go           # App entrypoint
└── README.md
```

## Roadmap

- Add support for more vulnerability types (e.g. CSRF, LFI)
- Export scan results to JSON
- Add terminal UI (TUI) mode using tview
- Add dry-run and interactive modes
- Implement object capability revocation (for Web3 use cases)

## Contributing

Contributions are welcome! Please fork the repo and submit a pull request, or open an issue for bugs or feature requests.

## Disclaimer

This tool is intended for educational and authorized security testing only. Unauthorized scanning or attacking websites you do not own or have permission to test is illegal and unethical.

## License

MIT License — see `LICENSE` file for details.