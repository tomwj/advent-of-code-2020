package main

import (
	"bufio"
	"fmt"
	"github.com/apex/log/handlers/cli"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/apex/log"
)

type PasswordEntry struct {
	MinCount int64
	MaxCount int64
	Letter   string
	Password string
}

type Slope struct {
	x int
	y int
}

func openFileLines(ctx log.Interface, fileName string) (err error, lines []string) {
	defer ctx.WithField("fileName", fileName).Debug("debug")
	file, err := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInFile := scanner.Text()
		lines = append(lines, lineInFile)
	}
	return err, lines
}

func sumTwoMultiply2020(accounts []int64) int64 {
	for _, n := range accounts {
		for _, m := range accounts {
			if n+m == 2020 {
				return n * m
			}
		}
	}
	return 0
}
func sumThreeMultiply2020(ctx log.Interface, accounts []int64) int64 {
	for _, n := range accounts {
		for _, m := range accounts {
			for _, o := range accounts {
				if n+m+o == 2020 {
					ctx.WithFields(log.Fields{
						"n": n,
						"m": m,
						"o": o,
					}).Debug("numbers used for result")
					return n * m * o
				}
			}
		}
	}
	return 0
}
func parseIntEntry(ctx log.Interface, s []string) []int64 {

	var ints []int64
	for i, line := range s {
		number, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			ctx.WithField("At line: ", i+1).Trace("parsing").Stop(&err)
		}
		ints = append(ints, number)

	}
	return ints
}

func parsePasswordEntry(ctx log.Interface, s string) PasswordEntry {

	// 18-20 q: xqqqwmqgtcqnqqxgsqcq
	ctx.WithField("parsing", s)
	entryArray := strings.Split(s, " ")
	ctx.WithField("splitting string ", entryArray).Debugf("%v", entryArray)
	minMax := strings.Split(entryArray[0], "-")

	minCount, _ := strconv.ParseInt(minMax[0], 10, 64)
	maxCount, _ := strconv.ParseInt(minMax[1], 10, 64)

	return PasswordEntry{
		MinCount: minCount,
		MaxCount: maxCount,
		Letter:   string(entryArray[1][0]),
		Password: entryArray[2],
	}
}

func passwordIsValidPart1(ctx log.Interface, password PasswordEntry) bool {
	ctx.Debugf("%v", password)
	var letterCount int64 = 0
	for _, letter := range password.Password {
		if string(letter) == password.Letter {
			letterCount++
		}
	}
	return (letterCount >= password.MinCount) && (letterCount <= password.MaxCount)
}

func passwordIsValidPart2(ctx log.Interface, password PasswordEntry) bool {
	isValid := (password.Letter == string(password.Password[password.MinCount-1])) !=
		(password.Letter == string(password.Password[password.MaxCount-1]))

	defer ctx.Debugf("Valid password %v %t", password, isValid)
	return isValid
}

func getColumn(ctx log.Interface, grid []string, columnNumber int) string {
	columnWrapAround := columnNumber % len(grid[0])
	var column string
	for _, row := range grid {
		column = column + string(row[columnWrapAround-1])
	}
	ctx.Debugf("column %v ", column)
	return column
}

func getValueInGrid(ctx *log.Entry, x int, y int, grid []string) interface{} {
	// Trim newline
	gridSize := len(grid[0])
	// 0 index x and y
	x--
	y--
	if x >= gridSize {
		ctx.Debugf("wrapping on grid %d, x %d after x %d", gridSize, x, (x % gridSize))
		x = x % gridSize
	}
	// Grid is an array of strings, each row is a string
	row := grid[y]
	value := string(row[x])
	ctx.Debugf("value %v at x: %d y: %d ", value, x, y)
	return value
}

func getCountTrees(ctx *log.Entry, slope Slope, grid []string) int {

	countOfTrees := 0
	for x, y := 1, 1; y <= len(grid); x, y = x+slope.x, y+slope.y {
		content := getValueInGrid(ctx, x, y, grid)
		if content == "#" {
			countOfTrees++
			ctx.Debugf("countOfTrees %d", countOfTrees)
		}
	}

	return countOfTrees
}
func parsePassports(ctx log.Interface, linesInFile []string) []string {
	var passes []string
	passIndex := 0
	passes = append(passes, "")
	for i, line := range linesInFile {
		if len(line) > 0 {
			passes[passIndex] += line + " "
			ctx.Debugf("Found line %s", line)
		}
		if len(line) == 0 {
			ctx.Debugf("Found empty line %s at %d, starting new pass", line, i+1)
			passes[passIndex] = strings.TrimSpace(passes[passIndex])
			passes = append(passes, "")
			passIndex++
		}
	}
	for _, pass := range passes {
		fmt.Printf("pass %v \n", pass)
	}
	return passes
}
func validatePassport(ctx log.Interface, passport string) bool {

	passport = strings.TrimSpace(passport)
	passwordFields := strings.Split(passport, " ")
	var fields [][]string
	for _, field := range passwordFields {
		fields = append(fields, strings.Split(field, ":"))
	}
	//fmt.Printf("passport fields %v \n", fields)
	requiredFields := []string{"byr", "cid", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}
	// Go through the given fields, and remove each one from the required list only if it is valid
	for _, field := range fields {

		switch field[0] {

		case "byr":
			//byr (Birth Year) - four digits; at least 1920 and at most 2002.
			year, _ := strconv.ParseInt(field[1], 10, 64)
			if year >= 1920 && year <= 2002 {

				ctx.Debugf("Valid birth year (byr) found %v", field[1])
				requiredFields = remove(requiredFields, field[0])
				ctx.Debugf("requireFields after %v", requiredFields)

			} else {
				ctx.Debugf("Invalid birth year (byr) found %v", field[1])
			}
		case "iyr":
			//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
			year, _ := strconv.ParseInt(field[1], 10, 64)
			if year >= 2010 && year <= 2020 {

				ctx.Debugf("Valid issue year (iyr) found %v", field[1])
				requiredFields = remove(requiredFields, field[0])
				ctx.Debugf("requireFields after %v", requiredFields)

			} else {
				ctx.Debugf("Invalid issue year (iyr) found %v", field[1])
			}
		case "eyr":
			//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
			year, _ := strconv.ParseInt(field[1], 10, 64)
			if year >= 2020 && year <= 2030 {
				requiredFields = remove(requiredFields, field[0])
			}
		case "hgt":
			//hgt (Height) - a number followed by either cm or in:
			//If cm, the number must be at least 150 and at most 193.
			//If in, the number must be at least 59 and at most 76.
			heightField := []rune(field[1])
			// Trim the last two chars
			height, _ := strconv.ParseInt(string(heightField[:len(heightField)-2]), 10, 64)
			if strings.Contains(field[1], "in") {
				if height >= 59 && height <= 76 {
					requiredFields = remove(requiredFields, field[0])
				}
			}
			if strings.Contains(field[1], "cm") {
				if height >= 150 && height <= 193 {
					requiredFields = remove(requiredFields, field[0])
				}

			}
		case "hcl":
			//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
			regexPattern := `(#[0-9a-f]{6})`
			matched, err := regexp.MatchString(regexPattern, field[1])
			if err != nil {
				ctx.Errorf("Regex for hcl wrong %s", regexPattern)
			}
			if matched {
				requiredFields = remove(requiredFields, field[0])
			}
		case "ecl":
			//	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			regexPattern := `(amb|blu|brn|gry|grn|hzl|oth)`
			matched, err := regexp.MatchString(regexPattern, field[1])
			if err != nil {
				ctx.Errorf("Regex for ecl wrong %s", regexPattern)
			}
			if matched {
				requiredFields = remove(requiredFields, field[0])
			}
		case "pid":
			//	pid (Passport ID) - a nine-digit number, including leading zeroes.
			regexPattern := `[0-9]`
			matched, err := regexp.MatchString(regexPattern, field[1])
			if err != nil {
				ctx.Errorf("Regex for ecl wrong %s", regexPattern)
			}
			if matched && len(field[1]) == 9 {
				requiredFields = remove(requiredFields, field[0])
			}
			//	cid (Country ID) - ignored, missing or not.
		}
	}
	fmt.Printf("Remaining fields %v on \npassport %s\n", requiredFields, passport)
	// All the fields were present
	if len(requiredFields) == 0 {
		return true
	}
	if len(requiredFields) == 1 {
		// Could be an Arctic pass
		for _, field := range requiredFields {
			if field == "cid" {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	functionName := runtime.FuncForPC(pc).Name()
	fn := functionName[strings.LastIndex(functionName, ".")+1:]
	ctx := log.WithFields(log.Fields{
		"project":  "advent-of-code",
		"file":     filename,
		"function": fn,
	})
	_, lines := openFileLines(ctx, "./data/day-1-input.txt")
	accountQuestion := parseIntEntry(ctx, lines)
	println("Day 1 Part 1 answer: ", sumTwoMultiply2020(accountQuestion))
	println("Day 1 Part 2 answer: ", sumThreeMultiply2020(ctx, accountQuestion))

	_, passwordDBDump := openFileLines(ctx, "./data/day-2-input.txt")
	var validCountPart1 = 0
	var validCountPart2 = 0
	for _, entry := range passwordDBDump {
		parsedPassword := parsePasswordEntry(ctx, entry)
		if passwordIsValidPart1(ctx, parsedPassword) {
			validCountPart1++
		}
		if passwordIsValidPart2(ctx, parsedPassword) {
			validCountPart2++
		}
	}
	println("Day 2 Part 1 answer: ", validCountPart1)
	println("Day 2 Part 2 answer: ", validCountPart2)

	_, linesInFile := openFileLines(ctx, "./data/day-3-trees.txt")
	slopeTest := Slope{3, 1}
	treeCount := getCountTrees(ctx, slopeTest, linesInFile)
	fmt.Printf("Day 3 Part 2 Trees answer: %d", treeCount)

	_, linesInFile = openFileLines(ctx, "./data/day-4-passports.txt")
	passportList := parsePassports(ctx, linesInFile)
	validPassports := 0
	for _, passport := range passportList {
		if validatePassport(ctx, passport) {
			validPassports++
		}
	}
	fmt.Printf("Day 4 Part 1 passports answer: %d", validPassports)
}
