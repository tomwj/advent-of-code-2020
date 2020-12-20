package main

import (
	"github.com/apex/log"
	"testing"
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
