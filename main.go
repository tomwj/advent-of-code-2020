package main

import (
	"bufio"
	"github.com/apex/log/handlers/cli"
	"os"
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

func getCountTrees(ctx *log.Entry, slope Slope, grid []string) interface{} {

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

}
