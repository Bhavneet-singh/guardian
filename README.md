# Guardian

Guardian is a small Go-based CLI tool for scanning NPM `package.json` files.
It reads dependencies and devDependencies, prints a summary, and supports both verbose and structured JSON output.

## Overview

The primary feature is the `scan` command, which parses a `package.json` file and returns details about the package name, version, dependency counts, and dependency listings.

This repository includes a Go CLI application under `guardian/guardian`, with a minimal scanner implementation in `guardian/guardian/internal/scanner`.

## Features

- Scan an NPM `package.json` file
- Output package metadata and dependency counts
- Optional verbose output for readable listing
- Optional JSON output for machine-readable results
- Includes placeholder `hello` and `version` commands for CLI extension

## Repository Layout

- `guardian/guardian/` - Go module and CLI app
  - `main.go` - entry point for the CLI
  - `cmd/` - Cobra command definitions
    - `scan.go` - `scan` command implementation
    - `hello.go` - sample `hello` command
    - `version.go` - sample `version` command
    - `root.go` - root command setup
  - `internal/scanner/npm.go` - package.json scanning logic
  - `go.mod` - Go module definition

## Requirements

- Go 1.25+

## Build

From the repository root:

```powershell
cd guardian/guardian
go build -o guardian.exe
```

## Usage

From `guardian/guardian`:

```powershell
./guardian.exe scan
```

Scan a specific `package.json` file:

```powershell
./guardian.exe scan --file ../some/path/package.json
```

Enable verbose output:

```powershell
./guardian.exe scan --verbose
```

Enable JSON output:

```powershell
./guardian.exe scan --json
```

## Examples

Scan the default `package.json` in the current directory:

```powershell
./guardian.exe scan
```

Scan a file at a custom path:

```powershell
./guardian.exe scan --file ../guardian/guardian/package.json
```

## Development

- Extend `cmd/scan.go` to add more scanner options.
- Add commands under `cmd/` with Cobra.
- Improve parsing in `internal/scanner/npm.go` to support additional package metadata.

## License

This project is currently unlicensed. Update the repository license as needed.
