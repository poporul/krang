package krang_test

import (
	"fmt"
	"testing"

	"github.com/poporul/krang"
)

const (
	OK  = "[32m✓[0m"
	BAD = "[31m✘[0m"
)

var data = map[string]int{
	"2 + 3 * 4 - 5 + 7 * 6 / 3 - 2 * 3 ^ 2 + ( 5 - 2 ) * 2": 11,
	"3 + 4 * 2 / ( 5 - 1 ) ^ 2 ^ 3 + 2":                     5,
	"   4 -(  1 +3) + 81 /9 ":                               9,
}

func buildLog(status, source string, expected, actual int) string {
	return fmt.Sprintf(
		"%s %s (expected: %d, actual: %d)", status, source, expected, actual,
	)
}

func TestEval(t *testing.T) {
	for source, expected := range data {
		actual := krang.Eval(source)

		if actual != expected {
			t.Error(buildLog(BAD, source, expected, actual))
		} else {
			t.Log(buildLog(OK, source, expected, actual))
		}
	}
}
