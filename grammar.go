package krang

import (
	"errors"
	"strconv"
	"strings"
)

var (
	TypeError    = errors.New("type mismatch")
	DivideByZero = errors.New("divide by zero")
	BadOperand   = errors.New("bad operand")
)

type Associativity string
const (
	Left Associativity = "left"
	Right Associativity = "right"
)

type Binary interface {
	Calculate(first, second Value) (*Value, error)
}

type Unary interface {
	Calculate(first Value) (*Value, error)
}

type Operator interface {
	Precedence() int
	Associativity() Associativity
}

type operator struct {
	precedence    int
	associativity Associativity
}

type Value struct {
	Val interface{}
}

func FromNum(n int) *Value {
	return &Value{Val: n}
}

func FromBool(b bool) *Value {
	return &Value{Val: b}
}

func FromString(s string) *Value {
	return &Value{Val: s}
}

func (v Value) AsInt() (int, error) {
	switch typed := v.Val.(type) {
	case int:
		return typed, nil
	case string:
		i, err := strconv.ParseInt(typed, 10, strconv.IntSize)
		if err != nil {
			return 0, BadOperand
		}
		return int(i), nil
	}
	return 0, TypeError
}

func (v Value) AsBool() (bool, error) {
	switch typed := v.Val.(type) {
	case bool:
		return typed, nil
	case string:
		if strings.EqualFold(typed, "true") {
			return true, nil
		} else if strings.EqualFold(typed, "false") {
			return false, nil
		} else {
			return false, BadOperand
		}
	}
	return false, TypeError
}

func (o operator) Precedence() int {
	return o.precedence
}

func (o operator) Associativity() Associativity {
	return o.associativity
}

type Grammar struct {
	Operators map[string]Operator
}
