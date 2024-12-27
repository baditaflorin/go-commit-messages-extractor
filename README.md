# Commit Messages Extractor

This Go application extracts all commit messages from a local Git repository and outputs them in a specified format.

## Features
- Extract commit messages from a Git repository.
- Supports output in:
    - Concatenated lines (one message per line).
    - JSON format.
- Optionally include or omit commit hashes in the output.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/baditaflorin/go-commit-messages-extractor.git
   ```

2. Navigate to the project directory:
   ```bash
   cd commit-messages-extractor
   ```

3. Build the application:
   ```bash
   go build -o commit-messages-extractor
   ```

## Usage

Run the application with the following flags:

- `--repo` (required): Path to the local Git repository.
- `--format` (optional): Output format (`line` or `json`). Default is `line`.
- `--hash` (optional): Include commit hashes in the output (`true` or `false`). Default is `true`.

### Examples

#### Output as Lines with Hashes
```bash
./commit-messages-extractor --repo=/path/to/local/repo --format=line --hash=true
```

#### Output as Lines without Hashes
```bash
./commit-messages-extractor --repo=/path/to/local/repo --format=line --hash=false
```

#### Output as JSON
```bash
./commit-messages-extractor --repo=/path/to/local/repo --format=json
```

## Dependencies
- [go-git](https://pkg.go.dev/github.com/go-git/go-git/v5)

Install dependencies with:
```bash
go mod tidy
```

## License

This project is licensed under the GNU GENERAL PUBLIC LICENSE. See the `LICENSE` file for details.
