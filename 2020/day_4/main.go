package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passport contains all the fields, only cid is optional
type Passport struct {
	fields map[string]string
}

// Validate is specific to solving part 2
func (p Passport) Validate() bool {
	numRequiredFields := len(p.fields)
	if numRequiredFields < 7 {
		return false
	} else if _, exists := p.fields["cid"]; numRequiredFields == 7 && exists {
		return false
	}

	birthYear, errBirth := strconv.Atoi(p.fields["byr"])
	issueYear, errIssue := strconv.Atoi(p.fields["iyr"])
	expirationYear, errExpiration := strconv.Atoi(p.fields["eyr"])

	if errBirth != nil || birthYear < 1920 || birthYear > 2002 {
		return false
	}
	if errIssue != nil || issueYear < 2010 || issueYear > 2020 {
		return false
	}
	if errExpiration != nil || expirationYear < 2020 || expirationYear > 2030 {
		return false
	}

	heightMatch, err := regexp.MatchString("^[0-9]{2,3}(cm|in)$", p.fields["hgt"])
	if !heightMatch || err != nil {
		return false
	}
	height := p.fields["hgt"]
	heightMeasurement := height[len(height)-2:]
	heightVal, _ := strconv.Atoi(height[:len(height)-2])
	if heightMeasurement == "cm" && (heightVal < 150 || heightVal > 193) {
		return false
	} else if heightMeasurement == "in" && (heightVal < 59 || heightVal > 76) {
		return false
	}

	hairColorMatch, err := regexp.MatchString("^#[0-9a-f]{6}$", p.fields["hcl"])
	if !hairColorMatch || err != nil {
		return false
	}

	eyeColorMatch, err := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", p.fields["ecl"])
	if !eyeColorMatch || err != nil {
		return false
	}

	passportIDMatch, err := regexp.MatchString("^[0-9]{9}$", p.fields["pid"])
	if !passportIDMatch || err != nil {
		return false
	}

	return true
}

//NewPassport is a factory function
func NewPassport() Passport {
	p := Passport{}
	p.fields = make(map[string]string)
	return p
}

func createPassports(inputFilename string) []Passport {
	f, _ := os.Open(inputFilename)
	scanner := bufio.NewScanner(f)

	passports := []Passport{}
	currPassport := NewPassport()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, currPassport)
			currPassport = NewPassport()
		} else {
			entries := strings.Split(line, " ")
			for _, entry := range entries {
				keyValuePair := strings.Split(entry, ":")
				field, value := keyValuePair[0], keyValuePair[1]
				currPassport.fields[field] = value
			}
		}
	}

	passports = append(passports, currPassport)

	return passports
}

func partTwo(inputFilename string) {
	res := 0
	passports := createPassports(inputFilename)
	for _, passport := range passports {
		if passport.Validate() {
			res++
		}
	}
	fmt.Println(res)
}

func main() {
	partTwo("input.txt")
}
