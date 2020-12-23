package krang

import "reflect"

func EqualityOperator() Operator {
	return &equalityOperator{operator{
		precedence:    0,
		associativity: Left,
	}}
}

type equalityOperator struct {
	operator
}

func (o equalityOperator) Calculate(first, second Value) (*Value, error) {
	if reflect.TypeOf(first.Val) == reflect.TypeOf(second.Val) {
		return FromBool(first.Val == second.Val), nil
	}
	f, err := first.AsBool()
	if err != nil {
		return nil, TypeError
	}
	s, err := second.AsBool()
	if err != nil {
		return nil, TypeError
	}
	return FromBool(f == s), nil
}

func LogicalBinary(precedence int, associativity Associativity, fn func(bool, bool) (bool, error)) Operator {
	return logicalOperator{
		operator: operator{
			precedence:    precedence,
			associativity: associativity,
		},
		fn: fn,
	}
}

type logicalOperator struct {
	operator
	fn func(first, second bool) (bool, error)
}

func asBool(first, second Value) (bool, bool, error) {
	f, err := first.AsBool()
	if err != nil {
		return false, false, err
	}
	s, err := second.AsBool()
	if err != nil {
		return false, false, err
	}
	return f, s, nil
}

func (o logicalOperator) Calculate(first, second Value) (*Value, error) {
	f, s, err := asBool(first, second)
	if err != nil {
		return nil, err
	}
	res, err := o.fn(f, s)
	if err != nil {
		return nil, err
	}
	return FromBool(res), nil
}

func LogicalGrammar() Grammar {
	eq := EqualityOperator()
	and := LogicalBinary(-1, "left", func(first bool, second bool) (bool, error) {
		return first && second, nil
	})
	or := LogicalBinary(-2, "left", func(first bool, second bool) (bool, error) {
		return first || second, nil
	})
	g := Grammar{Operators: map[string]Operator{
		"=":  eq,
		"==": eq,
		"&":  and,
		"&&": and,
		"|":  or,
		"||": or,
	},
	}
	return g
}