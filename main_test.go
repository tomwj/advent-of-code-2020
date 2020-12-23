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

func logInit() *log.Entry {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestSuite",
	})
	return ctx
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
func TestMultiSlopeCountTree(t *testing.T) {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)
	ctx := log.WithFields(log.Fields{
		"func": "TestCountTrees",
	})

	_, linesInFile := openFileLines(ctx, "./data/day-3-trees-example.txt")

	// Right 1, down 1.
	// Right 3, down 1. (This is the slope you already checked.)
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	productOfTreeCounts := 1
	for _, slope := range slopes {
		treeCount := getCountTrees(ctx, slope, linesInFile)
		productOfTreeCounts = productOfTreeCounts * treeCount
		ctx.Infof("Tree count: %d", treeCount)
	}
	ctx.Infof("Product of tree counts %d", productOfTreeCounts)
	if productOfTreeCounts != 336 {
		t.Errorf("wanted 336 got %v", productOfTreeCounts)
	}
}
func TestMultiSlopeCountTreeSolution(t *testing.T) {
	ctx := logInit()
	_, linesInFile := openFileLines(ctx, "./data/day-3-trees.txt")

	// Right 1, down 1.
	// Right 3, down 1. (This is the slope you already checked.)
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	productOfTreeCounts := 1
	for _, slope := range slopes {
		treeCount := getCountTrees(ctx, slope, linesInFile)
		productOfTreeCounts = productOfTreeCounts * treeCount
		ctx.Infof("Tree count: %d", treeCount)
	}
	ctx.Infof("Product of tree counts %d", productOfTreeCounts)
	want := 2224913600
	if productOfTreeCounts != want {
		t.Errorf("wanted %d got %v", want, productOfTreeCounts)
	}
}
func TestParsePassports(t *testing.T) {
	ctx := logInit()

	_, linesInFile := openFileLines(ctx, "./data/day-4-passports-example.txt")

	want := "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"
	parsedPassports := parsePassports(ctx, linesInFile)
	got := parsedPassports[2]
	if got != want {
		t.Errorf("wanted %s got %s", want, got)
	}
}
func TestValidatePassportEntry(t *testing.T) {
	ctx := logInit()

	_, linesInFile := openFileLines(ctx, "./data/day-4-passports-example.txt")

	parsedPassports := parsePassports(ctx, linesInFile)
	for _, test := range []struct {
		expected bool
		entry    int
	}{
		{true, 0},
		{false, 1},
		{true, 2},
		{false, 3},
	} {
		passportValidation := validatePassport(ctx, parsedPassports[test.entry])
		if passportValidation != test.expected {
			t.Errorf(
				"Expected %t got %t for %s",
				test.expected, passportValidation, parsedPassports[test.entry])
		}
	}
}
func TestValidatePassportEntryPt2(t *testing.T) {
	ctx := logInit()

	_, linesInFile := openFileLines(ctx, "./data/day-4-examples-pt2-invalid.txt")

	parsedPassportsInvalid := parsePassports(ctx, linesInFile)
	for _, passport := range parsedPassportsInvalid {
		passportValidation := validatePassport(ctx, passport)
		if passportValidation {
			t.Errorf("Expected invalid got %t", passportValidation)
		}
	}
	_, linesInFile = openFileLines(ctx, "./data/day-4-examples-pt2-valid.txt")

	parsedPassportsValid := parsePassports(ctx, linesInFile)
	for _, passport := range parsedPassportsValid {
		passportValidation := validatePassport(ctx, passport)
		if !passportValidation {
			t.Errorf("Expected valid got %t", passportValidation)
		}
	}
}
