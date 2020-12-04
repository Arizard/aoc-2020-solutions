package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/arizard/aoc-2020-solutions/pkg/arie"
)

var required = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var validRegex = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile(`^\d{4}$`),
	"iyr": regexp.MustCompile(`^\d{4}$`),
	"eyr": regexp.MustCompile(`^\d{4}$`),
	"hgt": regexp.MustCompile(`^\d+(cm|in)$`),
	"hcl": regexp.MustCompile(`^#[a-f0-9]{6}$`),
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
	"cid": regexp.MustCompile(`.*`),
}

func isValidFormat(field string, value string) bool {
	pattern, ok := validRegex[field]
	fmt.Println(field, value)
	if !ok {
		panic("no regex for field")
	}
	isMatch := pattern.MatchString(value)
	if !isMatch {
		fmt.Println(field, value, "does not conform")
		return false
	}

	switch field {
	case "byr":
		num, _ := strconv.Atoi(value)
		return num >= 1920 && num <= 2002
	case "iyr":
		num, _ := strconv.Atoi(value)
		return num >= 2010 && num <= 2020
	case "eyr":
		num, _ := strconv.Atoi(value)
		return num >= 2020 && num <= 2030
	case "hgt":

		if strings.Contains(value, "cm") {
			num, _ := strconv.Atoi(strings.Replace(value, "cm", "", 1))
			return num >= 150 && num <= 193
		}
		if strings.Contains(value, "in") {
			num, _ := strconv.Atoi(strings.Replace(value, "in", "", 1))
			return num >= 59 && num <= 76
		}
		return false
	case "hcl":
		return true
	case "ecl":
		return true
	case "pid":
		return true
	case "cid":
		return true
	}

	return false
}

func main() {
	lines := arie.ReadSTDINLines()
	passports := []map[string]string{}
	kvpair := regexp.MustCompile(`([a-z]{3}:#*\w+)`)
	passportIndex := 0
	for _, line := range lines {
		if passportIndex > len(passports)-1 {
			passports = append(passports, map[string]string{})
		}

		matches := kvpair.FindAllString(line, -1)

		if len(matches) == 0 {
			passportIndex++
			continue
		}

		for _, match := range matches {
			args := strings.Split(match, ":")
			key := args[0]
			val := args[1]

			passports[passportIndex][key] = val
		}
	}

	valid := 0

	for _, passport := range passports {
		hasRequired := true
		for _, req := range required {
			if _, ok := passport[req]; !ok {
				hasRequired = false
			}
		}
		hasValidFormat := true
		for field, value := range passport {
			isValid := isValidFormat(field, value)
			fmt.Println(isValid)
			if !isValid {
				hasValidFormat = false
			}
		}
		if hasRequired && hasValidFormat {
			valid++
		}
	}

	fmt.Println(valid)
}
