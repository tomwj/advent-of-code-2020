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
