package day4

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var neededFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func Run(input string) {
	fileContent, _ := ioutil.ReadFile(input)
	fileString := string(fileContent)

	passports := strings.Split(fileString, "\n\n")

	fmt.Printf("Challenge 1: %v passports are valid\n", challenge1(passports))
	fmt.Printf("Challenge 2: %v passports are valid\n", challenge2(passports))
}

func challenge1(passports []string) (validPassports int) {
	for _, passport := range passports {
		invalid := false
		r := strings.NewReplacer("\n", " ")
		passport = r.Replace(passport)
		passport = strings.TrimSpace(passport)
		fields := make(map[string]string)
		for _, field := range strings.Split(passport, " ") {
			fields[strings.Split(field, ":")[0]] = strings.Split(field, ":")[1]
		}
		for _, neededField := range neededFields {
			if _, ok := fields[neededField]; !ok {
				invalid = true
				break
			}
		}
		if !invalid {
			validPassports++
		}
	}
	return validPassports
}

func challenge2(passports []string) (validPassports int) {
	for _, passport := range passports {
		invalid := false
		r := strings.NewReplacer("\n", " ")
		passport = r.Replace(passport)
		passport = strings.TrimSpace(passport)
		fields := make(map[string]string)
		for _, field := range strings.Split(passport, " ") {
			fields[strings.Split(field, ":")[0]] = strings.Split(field, ":")[1]
		}
		for _, neededField := range neededFields {
			if _, ok := fields[neededField]; !ok {
				invalid = true
				break
			}
		}
		if invalid {
			continue
		}

		byr, _ := strconv.Atoi(fields["byr"])
		if !(byr >= 1920 && byr <= 2002) {
			continue
		}

		iyr, _ := strconv.Atoi(fields["iyr"])
		if !(iyr >= 2010 && iyr <= 2020) {
			continue
		}

		eyr, _ := strconv.Atoi(fields["eyr"])
		if !(eyr >= 2020 && eyr <= 2030) {
			continue
		}

		hgt, _ := strconv.Atoi(fields["hgt"][:len(fields["hgt"])-2])
		suffix := fields["hgt"][len(fields["hgt"])-2:]
		if suffix == "cm" {
			if !(hgt >= 150 && hgt <= 193) {
				continue
			}
		} else if suffix == "in" {
			if !(hgt >= 59 && hgt <= 76) {
				continue
			}
		} else {
			continue
		}

		if match, _ := regexp.MatchString("^#[0-9a-f]{6}$", fields["hcl"]); !match {
			continue
		}

		if match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", fields["ecl"]); !match {
			continue
		}

		if match, _ := regexp.MatchString("^\\d{9}$", fields["pid"]); !match {
			continue
		}
		validPassports++
	}
	return validPassports
}
