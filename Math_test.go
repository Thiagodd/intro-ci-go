package main

import "testing"

func Test_sum(t *testing.T) {
	total := Sum(10, 20)

	if total != 30 {
		t.Errorf("Result of sum if invalid: Result %d. Expetecd: %d", total, 30)
	}
}
