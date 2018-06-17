# tomate
> Simple pomodoro 🍅
[![GoDoc](https://godoc.org/github.com/threkk/tomate?status.svg)](https://godoc.org/github.com/threkk/tomate) [![Code Climate](https://codeclimate.com/github/threkk/tomate/badges/gpa.svg)](https://codeclimate.com/github/threkk/tomate) [![Go Report Card](https://goreportcard.com/badge/github.com/threkk/tomate)](https://goreportcard.com/report/github.com/threkk/tomate) [![Sourcegraph](https://sourcegraph.com/github.com/threkk/tomate/-/badge.svg)](https://sourcegraph.com/github.com/threkk/tomate?badge)

## Features 
Simple pomodoro timer which a few useful functions:
- **Duration**: Duration can be changed to any time wished instead of the standard 25 minutes.
- **Repeat**: Enables the automatic restart of the timer once it finishes.
- **Messages**: Prints a message at the end of the timer.
- **Quiet mode**: Does not display any type of interface, useful for scripting.
- **Stop/pause/resume**: Stop with `Ctrl+C`, pause/resume with `Ctrl+Z`.

## Install

```
go get github.com/threkk/tomate
```

## Examples
- **Start a timer of 25 minutes.**
```
$ tomate
```

- **Start a timer of 4 minutes and 30 seconds and repeat.**
```
$ tomate -duration 4m30s -quiet
```

- **Start a timer of 1 minute in quite mode.**
```
$ tomate -duration 1m -quiet
```

- **Start a timer of 1 hour	with a message.**
```
$ tomate -duration 1h -message "🍅"
```

## Usage
A duration is a optionally signed sequence of decimal numbers followed by an
unit. A unit is `h` for hours, `m` for minutes and `s` for seconds.

## Meta
- **Author:** Alberto Martinez de Murga ([@threkk](https://threkk.com))
- **License:** BSD-3. See `LICENSE` for more information.
- **Repository:** https://github.com/threkk/tomate
