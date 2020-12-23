package krang_test

import (
	"fmt"
	"strings"
	"testing"
	"unicode"

	"github.com/poporul/krang"
)

const (
	OK  = "[32m✓[0m"
	BAD = "[31m✘[0m"
)

var (
	math = map[string]int{
		"2 + 3 * 4 - 5 + 7 * 6 / 3 - 2 * 3 ^ 2 + ( 5 - 2 ) * 2": 11,
		"3 + 4 * 2 / ( 5 - 1 ) ^ 2 ^ 3 + 2":                     5,
		"   4 -(  1 +3) + 81 /9 ":                               9,
	}

	logic = map[string]bool{
		"   abc = abc ":      true,
		"cba= abc ":          false,
		" (d = d) = true":    true,
		" true =  (d = d) ":  true,
		" (d = d) = false":   false,
		" false =  (d = c) ": true,
		" d = d & c = c" : true,
		" d = d && c = b" : false,
		" d = c | c = c" : true,
		" d = c || c = b" : false,
	}

	extensible = map[string]interface{}{
		" to_upper(abc) = ABC ":      true,
		" to_upper(abc) = abc ":      false,
		" to_lower(abc) = abc ":      krang.InvalidExpression,
		" is_digit(b) | is_digit(3) & is_digit(5)": true,
		" is_digit(b) | is_digit(3) & is_digit(a)": false,
		" (is_digit(b) | is_digit(3)) & is_digit(5)": true,
		" (is_digit(b) | is_digit(3)) & is_digit(d)": false,
		"get_name(42) == 'John Doe'": true,
	}

	damaged = map[string]error{
		"1 + a":       krang.BadOperand,
		"10 .* ( 9 )": krang.BadOperand,
		"1 + + 2":     krang.InvalidExpression,
		"+ 1 + 2":     krang.InvalidExpression,
		"1 + 2 +":     krang.InvalidExpression,
	}
)

func buildLog(status, source string, expected, actual interface{}) string {
	return fmt.Sprintf(
		"%s %s (expected: %v, actual: %v)", status, source, expected, actual,
	)
}

func TestLogic(t *testing.T) {
	g := krang.LogicalGrammar()
	for source, expected := range logic {
		actual, err := g.Eval(source)
		if err != nil {
			t.Errorf(buildLog(BAD, source, "no error", err.Error()))
		}
		res, err := actual.AsBool()
		if err != nil {
			t.Errorf(buildLog(BAD, source, "no error", err.Error()))
		}
		if res != expected {
			t.Error(buildLog(BAD, source, expected, res))
		} else {
			t.Log(buildLog(OK, source, expected, res))
		}
	}
}

func TestExtensible(t *testing.T) {
	mygrammar := krang.LogicalGrammar()
	mygrammar.Operators["to_upper"] = krang.UnaryFunc(func(v interface{}) (i interface{}, e error) {
		if s, ok := v.(string); ok {
			return strings.ToUpper(s), nil
		}
		return nil, krang.BadOperand
	})
	mygrammar.Operators["is_digit"] = krang.UnaryFunc(func(v interface{}) (i interface{}, e error) {
		if s, ok := v.(string); ok {
			return len(s) == 1 && unicode.IsDigit(rune(s[0])), nil
		}
		return nil, krang.BadOperand
	})
	mygrammar.Operators["get_name"] = krang.UnaryFunc(func(v interface{}) (i interface{}, e error) {
		return "John Doe", nil
	})
	for source, expected := range extensible {
		actual, err := mygrammar.Eval(source)
		if err != nil {
			if err != expected {
				t.Errorf(buildLog(BAD, err.Error(), 0, 0))
			} else {
				t.Log(buildLog(OK, source, expected, err.Error()))
			}
		} else {
				if actual.Val != expected {
					t.Error(buildLog(BAD, source, expected, actual.Val))
				} else {
					t.Log(buildLog(OK, source, expected, actual.Val))
				}
			}
	}
}

func TestMath(t *testing.T) {
	g := krang.MathGrammar()
	for source, expected := range math {
		actual, err := g.Eval(source)
		if err != nil {
			t.Fatalf(buildLog(BAD, err.Error(), 0, 0))
		}
		num, err := actual.AsInt()
		if err != nil {
			t.Fatalf(buildLog(BAD, err.Error(), 0, 0))
		}
		if num != expected {
			t.Error(buildLog(BAD, source, expected, num))
		} else {
			t.Log(buildLog(OK, source, expected, num))
		}
	}
}

func TestInvalidInput(t *testing.T) {
	g := krang.MathGrammar()
	for input, expectedError := range damaged {
		_, err := g.Eval(input)

		if err != expectedError {
			t.Error(BAD, input)
		} else {
			t.Log(OK, input)
		}
	}
}
