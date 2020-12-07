package main

import (
	"github.com/alecthomas/kong"
	day1 "github.com/jugendhacker/adventofcode/2020/Day1"
	day7 "github.com/jugendhacker/adventofcode/2020/Day7"
)

type CLI struct {
	Day       int    `help:"Day to run" short:"d"`
	InputPath string `arg name:"path" help:"Path to input file" type:"path"`
}

func main() {
	var cli CLI
	kong.Parse(&cli)
	switch cli.Day {
	case 1:
		day1.Run(cli.InputPath)
	case 7:
		day7.Run(cli.InputPath)
	}
}
