package main

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/google/go-cmp/cmp"
	"testing"
	"time"
)

var accounts = []int64{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func TestSumTwoMultiply2020(t *testing.T) {

	got := sumTwoMultiply2020(accounts)
	if got != 514579 {
		t.Errorf("SumTwoMultiply = %d; want 514579", got)
	}
}

func TestSumThreeMultiply2020(t *testing.T) {
	ctx := log.WithFields(log.Fields{
		"tests": "Test",
	})
	got := sumThreeMultiply2020(ctx, accounts)
	if got != 241861950 {
		t.Errorf("SumThreeMultiply = %d; want 241861950", got)
	}
}

func TestParsePasswordEntry(t *testing.T) {
	log.SetHandler(cli.Default)
	ctx := log.WithFields(log.Fields{
		"timestamp": time.Now(),
	})
	passwordEntries := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}
	want := PasswordEntry{
		1,
		3,
		"a",
		"abcde",
	}
	got := parsePasswordEntry(ctx, passwordEntries[0])
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("MakeGatewayInfo() mismatch (-want +got):\n%s", diff)
	}
}
func TestIsPasswordValidPart1(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestCheckPasswordEntry",
	})
	validPassword := PasswordEntry{
		1,
		3,
		"a",
		"abcde",
	}
	invalidPassword := PasswordEntry{
		1,
		3,
		"b",
		"cdefg",
	}
	if !passwordIsValidPart1(ctx, validPassword) {
		t.Error("valid password returned false")
	}
	if passwordIsValidPart1(ctx, invalidPassword) {
		t.Error("invalid password returned true")
	}
}

func TestIsPasswordValidPart2(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestCheckPasswordEntry",
	})
	validPassword := PasswordEntry{1, 3, "a", "abcde"}
	invalidPassword1 := PasswordEntry{1, 3, "b", "cdefg"}
	invalidPassword2 := PasswordEntry{2, 9, "c", "ccccccccc"}
	if !passwordIsValidPart2(ctx, validPassword) {
		t.Error("valid password returned false")
	}
	if passwordIsValidPart2(ctx, invalidPassword1) {
		t.Error("invalid password returned true")
	}
	if passwordIsValidPart2(ctx, invalidPassword2) {
		t.Error("invalid password returned true")
	}
}
func TestGetColumn(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestCheckPasswordEntry",
	})

	err, linesInFile := openFileLines(ctx, "./data/day-3-trees-example.txt")
	got := linesInFile[0]

	if err != nil {
		t.Errorf("Error reading file %v", err)
	}
	want := "..##......."
	if got != want {
		t.Errorf("wanted %s got %s", want, got)
	}

	// This is the first column transposed
	wantColumn := ".#......##."
	gotColumn := getColumn(ctx, linesInFile, 1)
	if gotColumn != wantColumn {
		t.Errorf("wanted %s got %s", wantColumn, gotColumn)
	}
	wantColumn2 := "..#.#.##..#"
	gotColumn2 := getColumn(ctx, linesInFile, 2)
	if gotColumn != wantColumn {
		t.Errorf("wanted col 2 %s got %s", wantColumn2, gotColumn2)
	}
	wantColumn14 := "#..#.#..#.."
	gotColumn14 := getColumn(ctx, linesInFile, 14)
	if gotColumn14 != wantColumn14 {
		t.Errorf("wanted col 2 %s got %s", wantColumn14, gotColumn14)
	}
}

func TestGetFromLocation(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestGetValue",
	})

	_, linesInFile := openFileLines(ctx, "./data/day-3-trees-example.txt")

	want := "."
	got := getValueInGrid(ctx, 1, 1, linesInFile)
	if got != want {
		t.Errorf("wanted %s got %v", want, got)
	}
	want = "#"
	got = getValueInGrid(ctx, 11, 8, linesInFile)
	if got != want {
		t.Errorf("wanted %s got %v", want, got)
	}
	want = "#"
	got = getValueInGrid(ctx, 12, 10, linesInFile)
	if got != want {
		t.Errorf("wanted %s got %v", want, got)
	}
}

func TestCountTree(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestCountTrees",
	})

	_, linesInFile := openFileLines(ctx, "./data/day-3-trees-example.txt")

	want := 7

	slopeTest := Slope{3, 1}
	got := getCountTrees(ctx, slopeTest, linesInFile)
	if got != want {
		t.Errorf("wanted %v got %v", want, got)
	}
}
