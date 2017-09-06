package krang

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

var (
	InvalidInputError = errors.New("Invalid input")
)

var operators = map[string]Operator{
	"^": Operator{"^", 4, "right"},
	"*": Operator{"*", 3, "left"},
	"×": Operator{"*", 3, "left"},
	"/": Operator{"/", 3, "left"},
	"÷": Operator{"/", 3, "left"},
	"-": Operator{"-", 2, "left"},
	"+": Operator{"+", 2, "left"},
}

type Number struct {
	Value int
}

func NewNumber(number string) Number {
	result, _ := strconv.Atoi(number)
	return Number{result}
}

type Operator struct {
	Value         string
	Precedence    int
	Associativity string
}

func (o Operator) Calculate(first, second Number) Number {
	var result int

	switch o.Value {
	case "*":
		result = first.Value * second.Value
	case "/":
		result = first.Value / second.Value
	case "+":
		result = first.Value + second.Value
	case "-":
		result = first.Value - second.Value
	case "^":
		result = int(math.Pow(float64(first.Value), float64(second.Value)))
	}

	return Number{result}
}

func (o Operator) Gte(right Operator) bool {
	return o.Precedence >= right.Precedence
}

func (o Operator) String() string {
	return fmt.Sprintf("{%s}", o.Value)
}

type LeftBracket struct{}
type RightBracket struct{}

type Token interface{}

func Tokenize(value string) ([]Token, error) {
	var lastNumber string
	var result []Token
	var err error

	for _, c := range value {
		if c >= '0' && c <= '9' {
			lastNumber = lastNumber + string(c)
		} else {
			if len(lastNumber) > 0 {
				result = append(result, NewNumber(lastNumber))
				lastNumber = ""
			}

			if c == ' ' {
				continue
			}

			if c == '(' {
				result = append(result, LeftBracket{})
				continue
			}

			if c == ')' {
				result = append(result, RightBracket{})
				continue
			}

			if operator, ok := operators[string(c)]; ok {
				result = append(result, operator)
			} else {
				err = InvalidInputError
			}
		}
	}

	if len(lastNumber) > 0 {
		result = append(result, NewNumber(lastNumber))
	}

	return result, err
}
