# Agents.md

## Project Overview
`jsto` is a CLI tool written in Go that displays the current time in various time zones relative to Japan Standard Time (JST). It is built using the Cobra library.

## Directory Structure
- `cmd/`: Contains the implementation of the CLI commands.
- `main.go`: The entry point of the application.

## Commands

### `edt`
Displays the time in Eastern Daylight Time (EDT).
- **Offset**: UTC-4 (JST-13)
- **Location**: America/New_York

### `ist`
Displays the time in Indian Standard Time (IST).
- **Offset**: UTC+5:30 (JST-3:30)
- **Location**: Asia/Kolkata

### `pdt`
Displays the time in Pacific Daylight Time (PDT).
- **Offset**: UTC-7 (JST-15)
- **Location**: America/Los_Angeles

### `utc`
Displays the time in Coordinated Universal Time (UTC).
- **Offset**: UTC+0 (JST-9)

## Usage
Run the tool using `jsto [command]`.
Example: `jsto utc`
