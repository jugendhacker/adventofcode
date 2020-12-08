package main

import (
	"github.com/alecthomas/kong"
	day1 "github.com/jugendhacker/adventofcode/2020/Day1"
	day2 "github.com/jugendhacker/adventofcode/2020/Day2"
	day3 "github.com/jugendhacker/adventofcode/2020/Day3"
	day4 "github.com/jugendhacker/adventofcode/2020/Day4"
	day5 "github.com/jugendhacker/adventofcode/2020/Day5"
	day6 "github.com/jugendhacker/adventofcode/2020/Day6"
	day7 "github.com/jugendhacker/adventofcode/2020/Day7"
	day8 "github.com/jugendhacker/adventofcode/2020/Day8"
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
	case 2:
		day2.Run(cli.InputPath)
	case 3:
		day3.Run(cli.InputPath)
	case 4:
		day4.Run(cli.InputPath)
	case 5:
		day5.Run(cli.InputPath)
	case 6:
		day6.Run(cli.InputPath)
	case 7:
		day7.Run(cli.InputPath)
	case 8:
		day8.Run(cli.InputPath)
	}
}
