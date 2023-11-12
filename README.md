# Advent of Code 2023

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This repository contains solutions to the 2023 [Advent of Code](https://adventofcode.com/) challenges.

## About

Every year I use the Advent of Code challenges to teach myself a new programming language and build up my problem-solving skills.

This year, I've chosen the language **[Go](go.dev)**, and will be practising [test driven development (TDD)](https://en.wikipedia.org/wiki/Test-driven_development) throughout the month.


## How to use this project

### Run all tests

This is run by CI on all branches

```sh
$ go list -f '{{.Dir}}/...' -m | xargs go test
```

### Solve a single day's puzzle

```sh
$ go run ./src/<day directory>
```