[![Coverage Status](https://coveralls.io/repos/github/jugendhacker/adventofcode/badge.svg?branch=main)](https://coveralls.io/github/jugendhacker/adventofcode?branch=main)

**⚠️ Warning: may contain spoilers ⚠️**

These are my solutions to the AdventOfCode challenges written in Golang.

# About

This project is made to build every challenge solution into one single binary, so you could play around with it.

To do this there is one `main.go` file which only selects which challenge to run. Each challenge solution resides in it's own sub package together with a unit test based on the explenations on the AdeventOfCode website. The main file then only calls the `Run()` function on the right package and gives it the path passed as commandline argument.


# How to get

## Download

Every commit on the main branch is build and release via Github Actions.

To get the build results just go to the release tab. There you could find binarys for linux and windows.

## Experienced Gopher: build it yourself

Just build the whole module with

```bash
go build -v .
```

# Usage

All my solutions are compile into a single binary and you could run it with

```bash
adventofcode -d 1 /path/to/input
```

For more information see

```bash
adventofcode --help
```
