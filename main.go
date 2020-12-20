package main

import (
	"bufio"
	"github.com/apex/log/handlers/cli"
	"os"
	"strconv"

	"github.com/apex/log"
)

func openFile(ctx log.Interface, fileName string) (err error, lines []int64) {
	defer ctx.WithField("fileName", fileName).Trace("opening").Stop(&err)
	file, err := os.Open(fileName)

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		lineInFile := scanner.Text()
		number, err := strconv.ParseInt(lineInFile, 10, 32)
		if err != nil {
			ctx.WithField("At line: ", lineNumber).Trace("parsing").Stop(&err)
		}
		lines = append(lines, number)
		lineNumber++
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
					}).Trace("return")
					return n * m * o
				}
			}
		}
	}
	return 0
}
func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"day": "1",
	})
	accounts := []int64{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	println(sumTwoMultiply2020(accounts))
	_, accountQuestion := openFile(ctx, "./data/day-1-input.txt")
	println("Part 1 answer: ", sumTwoMultiply2020(accountQuestion))
	println("Part 2 answer: ", sumThreeMultiply2020(ctx, accountQuestion))
}
